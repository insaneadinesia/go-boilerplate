package rest

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"maps"

	"github.com/go-resty/resty/v2"
	"github.com/insaneadinesia/gobang/logger"
	"go.opentelemetry.io/otel/trace"
)

type RestClient interface {
	Get(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error)
}

type client struct {
	HTTPClient *resty.Client
}

func New(opt Option) RestClient {
	httpClient := resty.New().
		SetBaseURL(opt.Address).
		SetTimeout(opt.Timeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: opt.SkipTLS})

	// Register Request-Response Middleware
	httpClient.OnBeforeRequest(BeforeRequestMiddleware)
	httpClient.OnAfterResponse(AfterRequestMiddleware)

	return &client{
		HTTPClient: httpClient,
	}
}

func BeforeRequestMiddleware(c *resty.Client, req *resty.Request) (err error) {
	if logger.Log == nil {
		return
	}

	ctx := req.Context()

	method := req.Method
	url := fmt.Sprintf("%s%s", c.BaseURL, req.URL)
	reqHeader := req.Header
	reqBody := req.Body

	logger.Log.Info(ctx, fmt.Sprintf("[Rest Header] %s %s", method, url), reqHeader)
	logger.Log.Info(ctx, fmt.Sprintf("[Rest Request] %s %s", method, url), reqBody)

	return nil
}

func AfterRequestMiddleware(c *resty.Client, resp *resty.Response) (err error) {
	if logger.Log == nil {
		return
	}

	ctx := resp.Request.Context()

	method := resp.Request.Method
	url := resp.Request.URL
	respBody := resp.Body()

	statusCode := resp.StatusCode()
	logStatusCode := map[string]any{
		"status_code": statusCode,
	}

	logger.Log.Info(ctx, fmt.Sprintf("[Rest Response] %s %s", method, url), string(respBody), logStatusCode)

	return nil
}

func (c *client) Get(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error) {
	body, statusCode, err = c.call(ctx, http.MethodGet, path, headers, nil)
	return
}

func (c *client) call(ctx context.Context, method, path string, requestHeader http.Header, requestBody []byte) (body []byte, status int, err error) {
	r := c.setRequest(ctx, requestHeader)
	resp, err := r.SetContext(ctx).SetBody(requestBody).Execute(method, path)
	if err != nil {
		if logger.Log != nil {
			logger.Log.Error(ctx, fmt.Sprintf("call %s %s%s error", method, c.HTTPClient.BaseURL, path), err.Error())
		}
	}

	body = resp.Body()
	status = resp.StatusCode()

	return
}

func (c *client) setRequest(ctx context.Context, requestHeader http.Header) (req *resty.Request) {

	// repopulate header, because resty cannot read http.Header
	headers := make(map[string][]string)
	maps.Copy(headers, requestHeader)
	req = c.HTTPClient.R().SetHeaderMultiValues(headers)

	// inject traceparent header
	req.SetHeaders(PopulateTraceparentHeadersFromSpanContext(trace.SpanFromContext(ctx).SpanContext()))

	return
}
