package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokersUrl, config)
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
