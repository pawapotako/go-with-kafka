package repository

import (
	"go-with-kafka/internal/model"
	"go-with-kafka/internal/util"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type externalRepositoryDB struct {
	db *gorm.DB
}

type ExternalRepository interface {
	EventStreaming(entity model.UserRequest)
}

func NewExternalRepositoryDB(db *gorm.DB) ExternalRepository {
	return externalRepositoryDB{db}
}

func (r externalRepositoryDB) EventStreaming(entity model.UserRequest) {

	config := util.LoadConfig()
	conn := util.KafkaConn(config, "user")

	// Check topic if already exists or not
	if !util.IsTopicAlreadyExists(conn, "user") {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             "user",
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}

	data := func() []kafka.Message {
		user := entity

		// Convert into kafka.Message{}
		messages := make([]kafka.Message, 0)

		messages = append(messages, kafka.Message{
			Value: util.CompressToJsonBytes(&user),
		})

		return messages
	}()

	// Set timeout
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(data...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// Close connection
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
