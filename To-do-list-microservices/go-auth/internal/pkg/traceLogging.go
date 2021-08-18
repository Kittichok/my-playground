package pkg

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
)

var TransactionName string

func InitTrace() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://d37f710b5e854df1bd97269f3a8adaa5@o952322.ingest.sentry.io/5901713",
		TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			return sentry.SampledTrue
		}),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}

func StartSpan(ctx context.Context, operation string) *sentry.Span {
	return sentry.StartSpan(ctx, operation, sentry.TransactionName(TransactionName))
}

func FinishSpan(span *sentry.Span) {
	span.Finish()
}

func SetupTransacName(name string) {
	TransactionName = name
}