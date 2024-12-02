package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/hentan/final_project/internal/logger"
)

type KafkaProducer interface {
	SendMessage(message string) error
	Close()
}

type KafkaProducerImpl struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducerImpl, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V2_8_0_0
	log := logger.GetLogger()

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Info("Ошибка создания Kafka Producer", "брокеры:", brokers, "конфиг:", config, "ошибка: ", err)
		return nil, err
	}

	return &KafkaProducerImpl{
		producer: producer,
		topic:    topic,
	}, nil
}

func (kp *KafkaProducerImpl) SendMessage(message string) error {
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

func (kp *KafkaProducerImpl) Close() {
	log := logger.GetLogger()
	if err := kp.producer.Close(); err != nil {
		sErr := fmt.Sprint("ошибка закрытия producer", err)
		log.Error(sErr)
	}
}
