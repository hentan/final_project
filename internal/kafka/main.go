package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/hentan/final_project/internal/logger"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	log := logger.GetLogger()

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Info("Ошибка создания Kafka Producer", "брокеры:", brokers, "топик:", topic)
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (kp *KafkaProducer) SendMessage(message string) error {
	msg := &sarama.ProducerMessage{
		Topic: kp.topic,
		Value: sarama.StringEncoder(message),
	}

	log := logger.GetLogger()

	_, _, err := kp.producer.SendMessage(msg)
	werr := fmt.Errorf("ошибка отправки сообщения %w", err)
	sErr := fmt.Sprint(werr)
	if err != nil {
		log.Error(sErr)
		return err
	}
	return nil
}

func (kp *KafkaProducer) Close() {
	log := logger.GetLogger()
	if err := kp.producer.Close(); err != nil {
		sErr := fmt.Sprint("ошибка закрытия producer", err)
		log.Error(sErr)
	}
}
