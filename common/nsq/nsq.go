package nsq

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

func NewProducer(topic string) chan []byte {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("localhost:4150", config)
	if err != nil {
		log.Fatal("Ошибка при создании продюсера:", err)
	}
	producer.SetLoggerLevel(nsq.LogLevelMax)

	ch := make(chan []byte)
	go func() {
		for v := range ch {
			if err := producer.Publish(topic, v); err != nil {
				log.Fatal("Ошибка при отправке сообщения:", err)
			}
		}

		producer.Stop()
	}()

	return ch
}

func NewConsumer(topic string) func() {
	config := nsq.NewConfig()

	// Создание консьюмера
	consumer, err := nsq.NewConsumer(topic, "notification", config)
	if err != nil {
		fmt.Println("Ошибка при создании консьюмера:", err)
		return func() {}
	}
	consumer.SetLoggerLevel(nsq.LogLevelMax)

	go func() {

		// Обработчик сообщений
		consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
			fmt.Printf("Получено сообщение: %s\n", message.Body)
			message.Finish()
			return nil
		}))

		// Подключение к NSQD
		err = consumer.ConnectToNSQD("localhost:4150")
		if err != nil {
			fmt.Println("Ошибка при подключении к NSQD:", err)
			return
		}

		<-consumer.StopChan
	}()

	return consumer.Stop
}
