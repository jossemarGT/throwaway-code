package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/IBM/sarama"
	cli "github.com/urfave/cli/v3"
)

func createProducer(c string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.ClientID = "kafka-go-producer"
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer(strings.Split(c, ","), config)
	if err != nil {
		log.Fatal("Error creating the producer:", err)
	}

	return producer
}

func main() {
	var topic, connstring string
	var partition, times int64
	var debug bool

	cmd := &cli.Command{
		Name:  "producer",
		Usage: "producer [FLAGS] message",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "topic",
				Value:       "default",
				Destination: &topic,
				Sources:     cli.EnvVars("TOPIC"),
			},
			&cli.IntFlag{
				Name:        "partition",
				Value:       0,
				Destination: &partition,
			},
			&cli.IntFlag{
				Name:        "times",
				Value:       1,
				Destination: &times,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Hidden:      true,
				Destination: &debug,
			},
			&cli.StringFlag{
				Name:        "conn",
				Hidden:      true,
				Value:       "localhost:9092",
				Destination: &connstring,
				Sources:     cli.EnvVars("CONN_STRING"),
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if debug {
				sarama.Logger = log.New(os.Stderr, "sarama: ", log.LstdFlags)
			}

			msg := c.Args().Get(0)

			producer := createProducer(connstring)
			defer func() {
				if err := producer.Close(); err != nil {
					log.Fatalln(err)
				}
			}()

			for i := 0; i < int(times); i++ {
				part, off, err := producer.SendMessage(&sarama.ProducerMessage{
					Topic:     topic,
					Partition: int32(partition),
					Value:     sarama.StringEncoder(fmt.Sprintf("%s %d", msg, i)),
				})

				if err != nil {
					return err
				}

				log.Printf("Sent message at partition %d offset %d", part, off)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal("it's dead")
	}
}
