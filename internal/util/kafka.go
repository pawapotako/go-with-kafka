package util

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func KafkaConn(cfg Config, topic string) *kafka.Conn {
	url := fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)
	conn, err := kafka.DialLeader(context.Background(), "tcp", url, topic, 0)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func IsTopicAlreadyExists(conn *kafka.Conn, topic string) bool {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partitions {
		if p.Topic == topic {
			return true
		}
	}
	return false
}
