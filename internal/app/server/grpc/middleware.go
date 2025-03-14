package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/gobang/gotel"
	"github.com/insaneadinesia/gobang/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TracingMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		// Start span
		tracer := gotel.DefaultTracerProvider().Tracer("grpc-server")
		ctx, span := tracer.Start(ctx, fmt.Sprintf("GRPC %s", info.FullMethod))
		defer span.End()

		return handler(ctx, req)
	}
}

func LoggingMiddleware(container *container.Container) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if err != nil {
				logger.Log.Error(ctx, "GRPC Request Error", err.Error())
			}

			logger.Log.TDR(ctx)
		}()

		start := time.Now()

		cfg := container.Config
		ctxLogger := logger.Context{
			ServiceName:    cfg.AppName,
			ServiceVersion: cfg.AppVersion,
			ServicePort:    cfg.AppHTTPPort,
			ReqMethod:      "GRPC",
			ReqURI:         info.FullMethod,
			ReqBody:        req,
		}

		ctx = logger.InjectCtx(ctx, ctxLogger)
		resp, err = handler(ctx, req)

		duration := time.Since(start)
		statusCode := codes.OK
		if err != nil {
			statusCode = status.Code(err)
		}

		ctxLogger.RespCode = int(statusCode)
		ctxLogger.RespBody = resp
		ctxLogger.RespTime = duration.String()

		ctx = logger.InjectCtx(ctx, ctxLogger)

		return resp, err
	}
}
