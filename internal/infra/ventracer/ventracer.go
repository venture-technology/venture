package ventracer

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type Ventracer struct {
	TracerProvider *sdktrace.TracerProvider
}

func NewJaegerTracer(host, service string) (*Ventracer, error) {

	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(host),
		otlptracehttp.WithInsecure(),
	)

	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		log.Fatalf("failed to create OTLP trace exporter: %v", err)
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
		)),
	)

	otel.SetTracerProvider(tp)

	return &Ventracer{TracerProvider: tp}, nil

}

func (v *Ventracer) Start(ctx context.Context, name string) (trace.Span, context.Context) {
	tracer := otel.Tracer("default-tracer")

	ctx, span := tracer.Start(ctx, name)

	return span, ctx
}
