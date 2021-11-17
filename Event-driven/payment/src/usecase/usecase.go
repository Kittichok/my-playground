package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kittichok/event-driven/payment/src/db/repository"
	"github.com/kittichok/event-driven/payment/src/event"
	models "github.com/kittichok/event-driven/payment/src/event/models"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type IUseCase interface {
	Payment(spanCtx opentracing.SpanContext, msg string) error
}

type UseCase struct {
	repo  repository.IRepository
	event event.Event
}

func NewUseCase(repo repository.IRepository, event event.Event) IUseCase {
	return UseCase{repo, event}
}

func (c UseCase) Payment(spanCtx opentracing.SpanContext, msg string) error {
	serverSpan := opentracing.GlobalTracer().StartSpan("Payment", ext.RPCServerOption(spanCtx))
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, serverSpan)
	defer serverSpan.Finish()
	b := []byte(msg)
	var booking models.BookingSubmitBody
	err := json.Unmarshal(b, &booking)
	if err != nil {
		fmt.Errorf("event submit booking error: %v", err.Error())
	}

	//TODO call 3party payment service?
	booking.Booking.PaymentStatus = "paid"
	body, err := json.Marshal(booking.Booking)
	c.event.SubmitMessage(ctx, serverSpan.Context(), "PaymentSuccess", string(body))

	return nil
}
