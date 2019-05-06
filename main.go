package main

import (
	"github.com/AoneFiles/kafka/kafkaSample/consumer"
	"github.com/AoneFiles/kafka/kafkaSample/producer"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	start      = kingpin.Flag("start", "start consumer or producer").String()
	broker     = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
	partition  = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	message    = kingpin.Flag("message", "message to be sent by the producer").Default("Hello World !").String()
)

func main() {
	kingpin.Parse()

	if *start == "" {
		panic("Please provide what service you need to run in the --start flag, check --help for details")
	}
	switch *start {
	case "consumer":
		consumer.Consumer(*broker, *topic, *partition, *offsetType)
	case "producer":
		producer.Producer(*broker, *topic, *message)
	}
}
