package consumer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// Consumer creates a consumer for the kafka server
func Consumer(broker []string, topic string, partition string, offsetType int) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	master, err := sarama.NewConsumer(broker, config)
	if err != nil {
		panic(err)
	}
	defer master.Close()
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Println("Message received from Producer: ", string(msg.Value))
			}
		}
	}()
	<-doneCh
}
