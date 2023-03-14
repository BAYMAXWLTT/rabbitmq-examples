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

	body := `hello,world`
	err = ch.Publish(`logs`, ``, false, false, amqp.Publishing{
		ContentType: `text/plain`,
		Body:        []byte(body),
	})
	utils.FailOnError(err, utils.ErrPublishingMessage.Error())
	log.Printf("[x] Sent %s", body)
}
