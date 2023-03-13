package main

import (
	"log"
	"simple/utils"

	"github.com/streadway/amqp"
)

func main() {
	// connect to rabbitmq server
	conn, err := amqp.Dial(utils.MQURL)
	utils.FailOnError(err, `error connecting to rabbitmq server`)
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	utils.FailOnError(err, `error creating a channel`)
	defer ch.Close()

	q, err := ch.QueueDeclare(`hello`, false, false, false, false, nil)
	utils.FailOnError(err, `error declaring a queue`)

	msgs, err := ch.Consume(q.Name, ``, true, false, false, false, nil)
	utils.FailOnError(err, `error registering a consumer`)

	forever := make(chan interface{})
	go func() {
		for m := range msgs {
			log.Printf("Received a message: %s\n", string(m.Body))
		}
	}()
	log.Printf("Waiting for messages, to exit press CTRL+C\n")
	// block the receiver
	<-forever
}
