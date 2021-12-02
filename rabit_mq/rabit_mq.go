package rabit_mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"notes-api/repositories"
	"os"
)

type HandleFuncMsg func(repo *repositories.LogRepository, message string)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConnectChannel() (*amqp.Channel, amqp.Queue) {
	urlConnect := fmt.Sprintf("amqp://%s:%s@%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_URL"))
	conn, err := amqp.Dial(urlConnect)
	failOnError(err, "Failed to connect to RabbitMQ%s")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	queue, err1 := ch.QueueDeclare(
		os.Getenv("RABBITMQ_NAME"), // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	failOnError(err1, "Failed to declare a queue")
	return ch, queue
}

func PublicMessage(ch *amqp.Channel, q amqp.Queue, body string) {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key`
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		failOnError(err, "Failed to public  message")
	}
}

func StartConsume(ch *amqp.Channel, q amqp.Queue, logRepo *repositories.LogRepository, handleFunc HandleFuncMsg) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			handleFunc(logRepo, string(d.Body))
		}
	}()
	<-forever
}
