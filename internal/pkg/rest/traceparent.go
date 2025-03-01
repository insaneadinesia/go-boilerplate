package rest

import (
	"fmt"

	"go.opentelemetry.io/otel/trace"
)

const (
	// TraceparentHeader is the HTTP header for trace propagation.
	//
	// For backwards compatibility, this is currently an alias for
	// for ElasticTraceparentHeader, but the more specific constants
	// below should be preferred. In a future version this will be
	// replaced by the standard W3C header.
	TraceparentHeader = ElasticTraceparentHeader

	// ElasticTraceparentHeader is the legacy HTTP header for trace propagation,
	// maintained for backwards compatibility with older agents.
	ElasticTraceparentHeader = "Elastic-Apm-Traceparent"

	// W3CTraceparentHeader is the standard W3C Trace-Context HTTP
	// header for trace propagation.
	W3CTraceparentHeader = "traceparent"

	// TracestateHeader is the standard W3C Trace-Context HTTP header
	// for vendor-specific trace propagation.
	TracestateHeader = "tracestate"
)

// Format span context from open telemetry with W3C traceparent header format
// https://www.w3.org/TR/trace-context/#traceparent-header
func FormatTraceparentHeader(span trace.SpanContext) string {
	const version = 0
	return fmt.Sprintf("%02x-%s-%s-%s", version, span.TraceID(), span.SpanID(), span.TraceFlags().String())
}

func PopulateTraceparentHeadersFromSpanContext(span trace.SpanContext) (headers map[string]string) {
	headers = make(map[string]string)
	headerValue := FormatTraceparentHeader(span)

	headers[ElasticTraceparentHeader] = headerValue
	headers[W3CTraceparentHeader] = headerValue

	if tracestate := span.TraceState().String(); tracestate != "" {
		headers[TracestateHeader] = tracestate
	}

	return
}
