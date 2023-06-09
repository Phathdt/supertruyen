package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

func WrapTraceIdFromIncomingContext(ctx context.Context, name, spanName string) (context.Context, trace.Span) {
	md, _ := metadata.FromIncomingContext(ctx)
	traceIdString := md["x-trace-id"][0]

	traceId, err := trace.TraceIDFromHex(traceIdString)
	if err != nil {
		return otel.Tracer(name).Start(ctx, spanName)
	}

	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	})

	// Embedding span config into the context
	ctx = trace.ContextWithSpanContext(ctx, spanContext)

	return otel.Tracer(name).Start(ctx, spanName)
}

func AppendTraceIdToOutgoingContext(ctx context.Context, name, spanName string) (context.Context, trace.Span) {
	ctx, span := otel.Tracer(name).Start(ctx, spanName)

	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	return ctx, span
}

func StartTrace(ctx context.Context, name, spanName string) (context.Context, trace.Span) {
	return otel.Tracer(name).Start(ctx, spanName)
}
