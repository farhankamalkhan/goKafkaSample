package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// Producer produces an input for the kafka server
func Producer(broker []string, topic string, message string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
