package rabbitmq

import (
	"encoding/json"

	"github.com/42core-team/bot-client/fail"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	STATUS_PULLING  = "pulling"
	STATUS_BUILDING = "building"
	STATUS_RUNNING  = "running"
	STATUS_FINISHED = "finished"
	STATUS_ERROR    = "error"
)

type Status struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func SendStatus(status string) {
	if !RabbitMQEnabled() {
		return
	}

	body, err := json.Marshal(Status{Status: status})
	fail.OnError(err, "Failed to marshal status")

	publishStatus(body)
}

func SendErrorStatus(errmsg string) {
	if !RabbitMQEnabled() {
		return
	}

	body, err := json.Marshal(Status{Status: STATUS_ERROR, Error: &errmsg})
	fail.OnError(err, "Failed to marshal status")
	publishStatus(body)
}

func publishStatus(body []byte) {
	if !RabbitMQEnabled() {
		return
	}

	err := ch.Publish(
		"",           // exchange
		statusq.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        body,
		})
	fail.OnError(err, "Failed to publish a message")
}
