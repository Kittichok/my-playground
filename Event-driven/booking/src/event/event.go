package event

import (
	"context"
	"fmt"
	"log"
	"time"

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

func NewEventConnection() Event {
	conn, err := kafka.DialLeader(context.Background(), "tcp", server, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// }
	return Event{conn}
}

func NewConsumer() {
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
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		eventProcesser(string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

func (e Event) SubmitMessage(ctx context.Context, eventName string, msg string) {
	_, err := e.conn.WriteMessages(
		kafka.Message{
			Key:   []byte(eventName),
			Value: []byte(msg),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

func eventProcesser(key string, msg string) {
	//TODO product price change?
	//TODO payment result success or fail
	return
}
