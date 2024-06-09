package rabbitmq

import (
	"log"
	"os"

	"github.com/42core-team/bot-client/fail"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var ch *amqp.Channel
var statusq amqp.Queue
var logq amqp.Queue

func Init() {
	init_conn()
	init_channel()
	init_queues()
}

func Close() {
	ch.Close()
	conn.Close()
}

func init_conn() {
	var err error
	conn, err = amqp.Dial(os.Getenv("RABBITMQ_URL"))
	fail.OnError(err, "Failed to connect to RabbitMQ")
}

func init_channel() {
	var err error
	ch, err = conn.Channel()
	fail.OnError(err, "Failed to open a channel")
}

func init_queues() {
	var err error
	clientID, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		log.Fatalln("CLIENT_ID not set")
	}
	statusq, err = ch.QueueDeclare(
		"bot.status." + clientID, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	fail.OnError(err, "Failed to declare the status queue")

	logq, err = ch.QueueDeclare(
		"bot.log." + clientID, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		amqp.Table{
			"x-queue-type": "stream",
		}, // arguments
	)
	fail.OnError(err, "Failed to declare the log queue")
}
