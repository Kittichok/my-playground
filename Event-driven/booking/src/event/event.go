package event

import (
	"context"
	"fmt"
	"log"

	"github.com/kittichok/event-driven/booking/src/db/models"
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"github.com/opentracing/opentracing-go"
	kafka "github.com/segmentio/kafka-go"
)

type Event struct {
	conn *kafka.Conn
}

const (
	topic     = "booking"
	partition = 0
	server    = "kafka:9092"
	groupID   = "consumer-group-id"
)

type EventName string

const (
	BookingSubmit  EventName = "BookingSubmit"
	BookingUpdated EventName = "BookingUpdated"
	PaymentSuccess EventName = "PaymentSuccess"
	PaymentFail    EventName = "PaymentFail"
)

type BookingSubmitBody struct {
	Booking       models.Booking
	BookingDetail []models.BookingDetail
}

type BookingUpdateBody struct {
	Booking models.Booking
}

func NewEventConnection() Event {
	conn, err := kafka.DialLeader(context.Background(), "tcp", server, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return Event{conn}
}

func NewConsumer(repo repository.IRepository) {
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
		eventProcesser(spanCtx, string(m.Key), string(m.Value), repo)
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

func (e Event) SubmitMessage(ctx context.Context, spanCtx opentracing.SpanContext, eventName EventName, msg string) {
	//FIX broken pipe connection
	//FIX remove implement e Event
	conn, err := kafka.DialLeader(ctx, "tcp", server, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	e.conn = conn

	_, err = e.conn.WriteMessages(
		kafka.Message{
			Key:     []byte(eventName),
			Value:   []byte(msg),
			Headers: SpanCtxToKafkaMessageHeaders(spanCtx),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := e.conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func eventProcesser(spanCtx opentracing.SpanContext, key string, msg string, repo repository.IRepository) {
	//TODO product price change?
	//TODO payment result success or fail
	if key == string(BookingUpdated) {
		updateBooking(spanCtx, msg, repo)
	} else if key == string(PaymentSuccess) {
		updateBooking(spanCtx, msg, repo)
	}
	return
}

func SpanCtxToKafkaMessageHeaders(spanCtx opentracing.SpanContext) []kafka.Header {
	m := make(opentracing.TextMapCarrier)
	opentracing.GlobalTracer().Inject(spanCtx, opentracing.TextMap, m)
	headers := make([]kafka.Header, 0, len(m))

	if err := m.ForeachKey(func(key, val string) error {
		headers = append(headers, kafka.Header{
			Key:   key,
			Value: []byte(val),
		})
		return nil
	}); err != nil {
		return headers
	}
	return headers
}

func TextMapCarrierFromKafkaMessageHeaders(headers []kafka.Header) opentracing.TextMapCarrier {
	textMap := make(map[string]string, len(headers))
	for _, header := range headers {
		textMap[header.Key] = string(header.Value)
	}
	return opentracing.TextMapCarrier(textMap)
}
