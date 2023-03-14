package main

import (
	"PubSub/utils"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(utils.MQURL)
	utils.FailOnError(err, utils.ErrConnectingToServer.Error())

	ch, err := conn.Channel()
	utils.FailOnError(err, utils.ErrCreatingChannel.Error())

	err = ch.ExchangeDeclare(`logs`, `fanout`, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringExchange.Error())

	q, err := ch.QueueDeclare(``, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringQueue.Error())

	err = ch.QueueBind(q.Name, ``, `logs`, false, nil)
	utils.FailOnError(err, utils.ErrBindingQueue.Error())

	msgs, err := ch.Consume(q.Name, ``, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrRegisteringConsumer.Error())

	log.Printf("Waiting for message, to exit press CTRL+C")
	forever := make(chan interface{})
	go func() {
		for m := range msgs {
			log.Printf("[*] Receive message: %s", string(m.Body))
		}
	}()
	<-forever
}
