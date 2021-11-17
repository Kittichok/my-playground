package processor

import (
	"context"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"
	kafka "github.com/segmentio/kafka-go"

	"github.com/kittichok/event-driven/payment/src/event"
	"github.com/kittichok/event-driven/payment/src/usecase"
)

const (
	topic   = "booking"
	server  = "kafka:9092"
	groupID = "payment-group"
)

func NewConsumer(usecase usecase.IUseCase) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{server},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	ctx := context.Background()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		carrierFromKafkaHeaders := TextMapCarrierFromKafkaMessageHeaders(m.Headers)
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, carrierFromKafkaHeaders)

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		err = eventProcesser(spanCtx, string(m.Key), string(m.Value), usecase)
		if err != nil {
			log.Fatal("process event error", err)
		}
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

func eventProcesser(spanCtx opentracing.SpanContext, key string, msg string, usecase usecase.IUseCase) error {
	if key == string(event.BookingSubmit) {
		return usecase.Payment(spanCtx, msg)
	}
	return nil
}

func TextMapCarrierFromKafkaMessageHeaders(headers []kafka.Header) opentracing.TextMapCarrier {
	textMap := make(map[string]string, len(headers))
	for _, header := range headers {
		textMap[header.Key] = string(header.Value)
	}
	return opentracing.TextMapCarrier(textMap)
}
