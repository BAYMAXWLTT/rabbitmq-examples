package main

import (
	"Routing/utils"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(utils.MQURL)
	utils.FailOnError(err, utils.ErrConnectingToServer.Error())

	ch, err := conn.Channel()
	utils.FailOnError(err, utils.ErrCreatingChannel.Error())

	err = ch.ExchangeDeclare(`logs_topic`, `topic`, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringExchange.Error())

	q, err := ch.QueueDeclare(``, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringQueue.Error())

	topic := os.Args[1]
	// bind the queue to a specified routing key of the exchange
	err = ch.QueueBind(q.Name, topic, `logs_topic`, false, nil)
	utils.FailOnError(err, utils.ErrBindingQueue.Error())

	msgs, err := ch.Consume(q.Name, ``, true, false, false, false, nil)

	log.Printf("Waiting for messages, to exit press CTRL+C")
	forever := make(chan interface{})
	go func() {
		for m := range msgs {
			log.Printf("[*] Receives message: %s, topic: %s", string(m.Body), topic)
		}
	}()
	<-forever
}
