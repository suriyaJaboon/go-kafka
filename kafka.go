package kafka

import (
	"github.com/Shopify/sarama"
)

const BootstrapServer = "localhost:9092"

type Kafka interface {
	DeleteTopic(topic string) error
	Close() error
}

type kafka struct {
	ca sarama.ClusterAdmin
}

func NewDefaultKafkaANDConfig(servers ...string) (Kafka, error) {
	if len(servers) == 0 {
		servers = []string{BootstrapServer}
	}

	admin, err := sarama.NewClusterAdmin(servers, nil)
	if err != nil {
		return nil, err
	}

	return &kafka{ca: admin}, nil
}

func NewKafka(ca sarama.ClusterAdmin) Kafka {
	return kafka{ca: ca}
}

func (k kafka) Close() error {
	return k.ca.Close()
}

func (k kafka) DeleteTopic(topic string) error {
	if err := k.ca.DeleteTopic(topic); err != nil {
		return err
	}

	return nil
}
