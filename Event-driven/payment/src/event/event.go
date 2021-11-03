package event

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

type Event struct {
	conn *kafka.Conn
}

const (
	topic     = "booking"
	partition = 0
	server    = "kafka:9092"
)

type EventName string

const (
	BookingSubmit  EventName = "BookingSubmit"
	BookingUpdated EventName = "BookingUpdated"
	PaymentSuccess EventName = "PaymentSuccess"
	PaymentFail    EventName = "PaymentFail"
)

func NewEventConnection() Event {
	conn, err := kafka.DialLeader(context.Background(), "tcp", server, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return Event{conn}
}

func (e Event) SubmitMessage(ctx context.Context, eventName EventName, msg string) {
	//FIX broken pipe connection
	//FIX remove implement e Event
	conn, err := kafka.DialLeader(context.Background(), "tcp", server, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	e.conn = conn

	_, err = e.conn.WriteMessages(
		kafka.Message{
			Key:   []byte(eventName),
			Value: []byte(msg),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := e.conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
