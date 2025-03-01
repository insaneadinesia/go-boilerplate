package rest

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/apperror"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	pkgValidator "github.com/insaneadinesia/go-boilerplate/internal/pkg/validator"
	"github.com/insaneadinesia/gobang/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ResponseRecorder is a custom response writer to capture the response body
// Because the app need custom logger middleware, it copy the logic of body dump middleware in echo
// Refs: https://github.com/labstack/echo/blob/master/middleware/body_dump.go#L52
type ResponseRecorder struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	err := http.NewResponseController(w.ResponseWriter).Flush()
	if err != nil && errors.Is(err, http.ErrNotSupported) {
		panic(errors.New("response writer flushing is not supported"))
	}
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return http.NewResponseController(w.ResponseWriter).Hijack()
}

func (w *bodyDumpResponseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func SetupMiddleware(server *echo.Echo, container *container.Container) {
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-App-Token, X-Client-Id"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.Use(middleware.Recover())
	server.Use(LoggingMiddleware(container))

	server.HTTPErrorHandler = ErrorHandler()
	server.Validator = &DataValidator{ValidatorData: pkgValidator.SetupValidator()}
}

func LoggingMiddleware(container *container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if SkipLoggerMiddleware(c.Path()) {
				return next(c)
			}

			start := time.Now()

			cfg := container.Config
			ctxLogger := logger.Context{
				ServiceName:    cfg.AppName,
				ServiceVersion: cfg.AppVersion,
				ServicePort:    cfg.AppHTTPPort,
				ReqMethod:      c.Request().Method,
				ReqURI:         c.Request().URL.String(),
				ReqHeader:      c.Request().Header,
			}

			// Request
			reqBody := []byte{}
			if c.Request().Body != nil { // Read
				reqBody, _ = io.ReadAll(c.Request().Body)
				ctxLogger.ReqBody = string(reqBody)
			}
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset

			// Inject Logger Context To Original Context
			ctx := logger.InjectCtx(c.Request().Context(), ctxLogger)
			c.SetRequest(c.Request().WithContext(ctx))

			// Response
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer

			if err = next(c); err != nil {
				c.Error(err)
			}

			// Calculate execution time
			execTime := time.Since(start).String()

			// Logging Response
			ctxLogger.RespCode = c.Response().Status
			ctxLogger.RespBody = string(resBody.Bytes())
			ctxLogger.RespTime = execTime

			ctx = logger.InjectCtx(ctx, ctxLogger)
			logger.Log.TDR(ctx)

			return
		}
	}
}

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		// Related Issue: https://github.com/labstack/echo/issues/1948
		if c.Response().Committed {
			return
		}

		resp := response.ErrorResponse{
			ErrorCode: constants.CODE_GENERAL_ERROR,
			DefaultResponse: response.DefaultResponse{
				Message: err.Error(),
			},
		}

		status := http.StatusBadRequest

		if ae, ok := err.(*apperror.ApplicationError); ok {
			status = ae.Status
			resp.Message = ae.Message
			resp.ErrorCode = ae.ErrorCode
		} else if ae, ok := err.(*echo.HTTPError); ok {
			status = ae.Code
		} else if _, ok := err.(runtime.Error); ok {
			status = http.StatusInternalServerError
		}

		err = c.JSON(status, resp)
	}
}

func SkipLoggerMiddleware(path string) bool {
	switch path {
	case "/health":
		return true
	}

	switch true {
	case strings.Contains(path, "/swagger"):
		return true
	}

	return false
}
