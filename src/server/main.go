package main

import (
	"github.com/gookit/event"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	_ "net/http/pprof"
	"picket/src/cmd"
	config2 "picket/src/config"
)

func main() {
	defer func() {
		err := event.CloseWait()
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	config, err := config2.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("fail to load config")
	}
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Error().Err(err).Send()
	}
	otel.SetTracerProvider(tp)
	root := cmd.Root(*config)

	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Msg("fail to execute root command")
	}

}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("picket-backend"),
			attribute.String("environment", "development"),
			attribute.Int64("ID", 1),
		)),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
}
