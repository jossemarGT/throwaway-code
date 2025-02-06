package main

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/IBM/sarama"
)

func main() {

	conn, ok := os.LookupEnv("CONN_STRING")
	if !ok || len(conn) == 0 {
		log.Fatal("Unable to find connection string on environment")
	}

	topic, ok := os.LookupEnv("TOPIC")
	if !ok || len(topic) == 0 {
		topic = "default"
	}

	debug, ok := os.LookupEnv("DEBUG")
	if ok && debug == "true" {
		sarama.Logger = log.New(os.Stderr, "sarama: ", log.LstdFlags)
	}

	config := sarama.NewConfig()
	config.ClientID = "kafka-go-consumer"
	consumer, err := sarama.NewConsumer(strings.Split(conn, ","), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	pconsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-pconsumer.Messages():
			log.Printf("[%s] %d %s", msg.Timestamp, msg.Partition, msg.Value)
		case <-signals:
			return
		}
	}
}
