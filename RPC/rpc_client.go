package main

import (
	"log"
	"os"
	"rpc/utils"

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

	corrID := utils.RandomString(32)
	q, err := ch.QueueDeclare(``, false, false, true, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringQueue.Error())

	n := os.Args[1]
	ch.Publish(``, `rpc_queue`, false, false, amqp.Publishing{
		ContentType:   `text/plain`,
		CorrelationId: corrID,
		ReplyTo:       q.Name,
		Body:          []byte(n),
	})
	msgs, err := ch.Consume(q.Name, ``, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrRegisteringConsumer.Error())

	for m := range msgs {
		if m.CorrelationId == corrID {
			ans := string(m.Body)
			log.Printf("ans: %s", ans)
			break
		}
	}
}
