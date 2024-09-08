package rabbitmq

import (
	"bufio"
	"io"
	"log"

	"github.com/42core-team/bot-client/fail"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StreamLogs(pipe io.ReadCloser) {
	if !RabbitMQEnabled() {
		return
	}

	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {
		line := scanner.Text()
		log.Print(line)

		err := ch.Publish(
			"",        // exchange
			logq.Name, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(line),
			})
		fail.OnError(err, "Failed to publish a message")
	}
}
