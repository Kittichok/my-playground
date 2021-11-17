package server

import (
	"fmt"
	"io"

	"github.com/kittichok/event-driven/booking/src/controllers"
	"github.com/kittichok/event-driven/booking/src/db"
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"github.com/kittichok/event-driven/booking/src/event"
	"github.com/kittichok/event-driven/booking/src/usecase"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"gorm.io/driver/sqlite"
)

func Init() {
	d := sqlite.Open("test.db")
	db.ConnectDataBase(d)
	_, closer := initJaeger("booking-service")
	defer closer.Close()
	rep := repository.NewRepository(db.DB)
	e := event.NewEventConnection()
	u := usecase.NewUseCase(rep, e)
	c := controllers.NewController(u)
	r := SetupRouter(c)

	go event.NewConsumer(rep)

	//TODO use viper get env
	port := "4000"
	r.Run(":" + port)

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
