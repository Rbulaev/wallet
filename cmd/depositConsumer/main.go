package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"wallet/internal/binding"
	"wallet/internal/model"
	"wallet/pkg/postgresql"
	"wallet/pkg/rabbitmq"
)

func main() {
	chName := "deposit"

	conn, err := rabbitmq.NewRabbitMQConn()
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"deposit",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, fmt.Sprintf("Failed to declare AMQP queue %s", chName))

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, fmt.Sprintf("Failed to register AMQP consumer %s", chName))

	postgresql, err := postgresql.Init()

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var body binding.Deposit
			err := json.Unmarshal(d.Body, body)
			handleError(err, fmt.Sprintf("Failed to marshal json %v to AMQP channel %s", body, chName))

			postgresql.First(&model.Wallet, d.Body).Update("amount", amount)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func handleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
