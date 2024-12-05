package kafka

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
	"github.com/theparthshira/go-image-search/elasticsearch"
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

func ConsumeImageTagData(client *elastic.Client) {
	brokers := []string{"localhost:9092"}
	topic := "imagetags"

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err)
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Failed to get partitions: ", err)
	}

	fmt.Println("Consuming...")

	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consuming partition %d: %v", partition, err)
		}
		defer partitionConsumer.Close()

		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n",
					msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

				str := string(msg.Value)
				trimmedStr := strings.Trim(str, "[]")

				tagArr := strings.Split(trimmedStr, ", ")

				for _, tag := range tagArr {
					elasticsearch.IndexElasticData(client, elasticsearch.PhotoTag{
						Tag: tag,
						Id:  string(msg.Key),
					})
				}

			}
		}(partitionConsumer)
	}

	select {}
}
