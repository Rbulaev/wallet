package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aftemark/wallet/internal/binding"
	"github.com/aftemark/wallet/pkg/rabbitmq"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
)

func main() {
	e := echo.New()

	e.POST("/deposit", deposit)
	e.POST("/transfer", transfer)

	e.Logger.Fatal(e.Start(":8080"))
}

func deposit(c echo.Context) error {
	var (
		chName  = "deposit"
		binding = new(binding.Deposit)
	)
	if err := c.Bind(&binding); err != nil {
		return err
	}

	if err := amqpHandler(binding, chName); err != nil {
		return err
	}

	return c.String(
		http.StatusOK,
		fmt.Sprintf("%#v was pushed to AMQP channel %s", binding, chName),
	)
}

func transfer(c echo.Context) error {
	var (
		chName  = "transfer"
		binding = new(binding.Transfer)
	)
	if err := c.Bind(&binding); err != nil {
		return err
	}

	if err := amqpHandler(binding, chName); err != nil {
		return err
	}

	return c.String(
		http.StatusOK,
		fmt.Sprintf("%#v was pushed to AMQP channel %s", binding, chName),
	)
}

func amqpHandler(binding interface{}, chName string) error {
	conn, err := rabbitmq.NewRabbitMQConn()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Failed to connect to RabbitMQ",
		)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Failed to connect to open a channel",
		)
	}
	defer ch.Close()

	if _, err := ch.QueueDeclare(
		chName,
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to declare AMQP channel %s", chName),
		)
	}

	body, err := json.Marshal(binding)
	if err != nil {
		return err
	}

	if err = ch.Publish(
		"",
		chName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to publish to AMQP channel %s: %#v", chName, body),
		)
	}

	return nil
}
