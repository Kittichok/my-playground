package server

import (
	"fmt"
	"io"

	"github.com/kittichok/event-driven/payment/src/controllers"
	"github.com/kittichok/event-driven/payment/src/db"
	"github.com/kittichok/event-driven/payment/src/db/repository"
	"github.com/kittichok/event-driven/payment/src/event"
	"github.com/kittichok/event-driven/payment/src/event/processor"
	"github.com/kittichok/event-driven/payment/src/usecase"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"gorm.io/driver/sqlite"
)

func Init() {
	d := sqlite.Open("test.db")
	db.ConnectDataBase(d)
	_, closer := initJaeger("payment-service")
	defer closer.Close()
	rep := repository.NewRepository(db.DB)
	e := event.NewEventConnection()
	u := usecase.NewUseCase(rep, e)
	c := controllers.NewController(u)
	r := SetupRouter(c)

	go processor.NewConsumer(u)

	//TODO use viper get env
	port := "4001"
	r.Run(":" + port)
}

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	_, err := cfg.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("cannot parse Jaeger env vars %v", err))
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
