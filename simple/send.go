package main

import (
	"log"

	"github.com/streadway/amqp"
	"simple/utils"
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

	// declare a queue for us to send to
	q, err := ch.QueueDeclare(`hello`, false, false, false, false, nil)
	utils.FailOnError(err, `error declaring a queue`)

	body := `hello,world`
	err = ch.Publish(``, q.Name, false, false, amqp.Publishing{
		ContentType: `text/plain`,
		Body:        []byte(body),
	})
	utils.FailOnError(err, `error publishing a message`)
	log.Printf("[x] Sent %s\n", body)
}
