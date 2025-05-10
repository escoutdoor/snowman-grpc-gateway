package tracing

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func Start(ctx context.Context, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	caller := functionCallerName(1)

	return tracer.Start(ctx, caller, opts...)
}

func StartWithName(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName, opts...)
}

func RecordError(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

func functionCallerName(depth int) string {
	fpcs := make([]uintptr, depth)

	n := runtime.Callers(3, fpcs)
	if n == 0 {
		panic("no caller")
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		panic("nil caller")
	}

	return caller.Name()
}
