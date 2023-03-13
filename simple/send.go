package main

import (
	"log"

	"github.com/streadway/amqp"
)

const MQURL = `amqp://guest:guest@127.0.0.1:5672/`

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}

func main() {
	// connect to rabbitmq server
	conn, err := amqp.Dial(MQURL)
	failOnError(err, `error connecting to rabbitmq server`)
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	failOnError(err, `error creating a channel`)
	defer ch.Close()

	// declare a queue for us to send to
	q, err := ch.QueueDeclare(`hello`, false, false, false, false, nil)
	failOnError(err, `error declaring a queue`)

	body := `hello,world`
	err = ch.Publish(``, q.Name, false, false, amqp.Publishing{
		ContentType: `text/plain`,
		Body:        []byte(body),
	})
	failOnError(err, `error publishing a message`)
	log.Printf("[x] Sent %s\n", body)
}
