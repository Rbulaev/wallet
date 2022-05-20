package rabbitmq

import (
	"github.com/streadway/amqp"
)

func NewRabbitMQConn() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@localhost:5672/")
}
