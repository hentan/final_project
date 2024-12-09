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
	producer sarama.AsyncProducer
	topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducerImpl, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V2_8_0_0
	log := logger.GetLogger()

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Info("Ошибка создания Kafka Producer", "брокеры:", brokers, "конфиг:", config, "ошибка: ", err)
		return nil, err
	}

	go func() {
		for {
			select {
			case msg := <-producer.Successes():
				log.Info("Сообщение успешно отправлено", "сообщение:", msg)
			case err := <-producer.Errors():
				log.Error(fmt.Sprintf("Ошибка отправки сообщения: %v", err))
			}
		}
	}()

	return &KafkaProducerImpl{
		producer: producer,
		topic:    topic,
	}, nil
}

func (kp *KafkaProducerImpl) SendMessage(message string) error {
	log := logger.GetLogger()
	msg := &sarama.ProducerMessage{
		Topic: kp.topic,
		Value: sarama.StringEncoder(message),
	}
	log.Info("Sending message to Kafka", "message", message)
	kp.producer.Input() <- msg

	return nil
}

func (kp *KafkaProducerImpl) Close() {
	log := logger.GetLogger()
	if err := kp.producer.Close(); err != nil {
		sErr := fmt.Sprint("ошибка закрытия producer", err)
		log.Error(sErr)
	}
}
