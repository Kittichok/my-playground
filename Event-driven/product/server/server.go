package server

import (
	"fmt"
	"io"

	"go.uber.org/dig"

	"github.com/kittichok/event-driven/product/controllers"
	"github.com/kittichok/event-driven/product/db"
	"github.com/kittichok/event-driven/product/db/repository"
	"github.com/kittichok/event-driven/product/usecase"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

type ServerService struct {
	dig.In

	Controller controllers.IController
}

func Init() {
	_, closer := initJaeger("product-service")
	defer closer.Close()
	container := dig.New()
	// container.Provide(event.NewEventService)
	container.Provide(db.NewSqliteConnection)
	container.Provide(repository.NewProductRepository)
	container.Provide(usecase.NewProductUseCase)
	container.Provide(controllers.NewController)

	httpServ := func(s ServerService) {
		r := SetupRouter(s.Controller)
		port := "5001"
		r.Run(":" + port)
	}
	if err := container.Invoke(httpServ); err != nil {
		panic(err)
	}
}

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	_, err := cfg.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("cannot parse Jaeger env vars %v", err))
	}

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
