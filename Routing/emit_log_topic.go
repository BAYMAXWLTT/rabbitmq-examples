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

	topic := os.Args[1]
	body := os.Args[2]
	err = ch.Publish(`logs_topic`, topic, false, false, amqp.Publishing{
		ContentType: `text/plain`,
		Body:        []byte(body),
	})
	utils.FailOnError(err, utils.ErrPublishingMessage.Error())

	log.Printf("[x] Sent %s, topic: %s", body, topic)
}
