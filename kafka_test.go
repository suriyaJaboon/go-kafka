package kafka

import (
	"github.com/Shopify/sarama"
	"testing"
)

func TestNewKafka(t *testing.T) {
	seedBroker := sarama.NewMockBroker(t, 1)
	defer seedBroker.Close()

	const tpc = "my_topic"

	seedBroker.SetHandlerByMap(map[string]sarama.MockResponse{
		"MetadataRequest": sarama.NewMockMetadataResponse(t).
			SetController(seedBroker.BrokerID()).
			SetBroker(seedBroker.Addr(), seedBroker.BrokerID()).
			SetLeader(tpc, 0, seedBroker.BrokerID()),
		"DeleteTopicsRequest": sarama.NewMockDeleteTopicsResponse(t),
	})

	config := sarama.NewConfig()
	//config.Version = sarama.V0_10_2_0 // Default from repository
	config.Version = sarama.V2_3_0_0

	admin, err := sarama.NewClusterAdmin([]string{seedBroker.Addr()}, config)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = admin.Close(); err != nil {
			t.Errorf("close cluster admin: %v", err)
		}
	}()

	t.Run("New default kafka and config", func(t *testing.T) {
		_, err = NewDefaultKafkaANDConfig()
		if err == nil {
			t.Errorf("New default kafka and config error: %v", err)
		}
	})
	t.Run("New default kafka and config close brokers", func(t *testing.T) {
		var kf Kafka
		kf, err = NewDefaultKafkaANDConfig(seedBroker.Addr())
		if err != nil {
			t.Errorf("New default kafka and config error: %v", err)
		}

		err = kf.Close()
		if err != nil {
			t.Errorf("New default kafka and config error: %v", err)
		}

	})
	t.Run("New default kafka and config success", func(t *testing.T) {
		_, err = NewDefaultKafkaANDConfig()
		if err == nil {
			t.Errorf("New default kafka and config error: %v", err)
		}
	})

	kf := NewKafka(admin)
	if kf == nil {
		t.Errorf("New kafka is nil")
	}

	t.Run("DeleteTopic Success", func(t *testing.T) {
		err = kf.DeleteTopic(tpc)
		if err != nil {
			t.Errorf("Delete topic name: %s %v", tpc, err)
		}
	})
	t.Run("DeleteTopic Empty", func(t *testing.T) {
		err = kf.DeleteTopic("")
		if err == nil {
			t.Errorf("Delete topic empty name is not nil")
		}
	})
}
