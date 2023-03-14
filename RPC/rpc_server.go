package main

import (
	"log"
	"rpc/utils"
	"strconv"

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

	utils.FailOnError(ch.Qos(1, 0, false), `error setting Qos`)

	_, err = ch.QueueDeclare(`rpc_queue`, false, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringQueue.Error())

	msgs, err := ch.Consume(`rpc_queue`, ``, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrRegisteringConsumer.Error())

	forever := make(chan interface{})
	go func() {
		for m := range msgs {
			n, err := strconv.Atoi(string(m.Body))
			log.Printf("Receive. n=%d", n)
			utils.FailOnError(err, `failed to convert body to integer`)
			ans := utils.Fibo(n)
			err = ch.Publish(``, m.ReplyTo, false, false, amqp.Publishing{
				ContentType:   `text/plain`,
				CorrelationId: m.CorrelationId,
				Body:          []byte(strconv.Itoa(ans)),
			})
			utils.FailOnError(err, utils.ErrPublishingMessage.Error())
		}
	}()
	log.Printf("Awaiting RPC requests...")
	<-forever
}
