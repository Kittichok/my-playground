package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kittichok/event-driven/booking/src/db/models"
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func updateBooking(spanCtx opentracing.SpanContext, msg string, repo repository.IRepository) {
	serverSpan := opentracing.GlobalTracer().StartSpan("UpdateBooking", ext.RPCServerOption(spanCtx))
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, serverSpan)
	defer serverSpan.Finish()
	b := []byte(msg)
	var booking models.Booking
	err := json.Unmarshal(b, &booking)
	if err != nil {
		fmt.Errorf("event update booking error: %v", err.Error())
	}
	err = repo.UpdateBooking(booking)
	if err != nil {
		fmt.Errorf("event update booking error: %v", err.Error())
	}
}
