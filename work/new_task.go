package main

import (
	"log"
	"os"
	"strings"
	"work/utils"

	"github.com/streadway/amqp"
)

func bodyFrom(args []string) string {
	if len(args) < 2 || os.Args[1] == `` {
		return `Hello`
	}
	return strings.Join(args[1:], " ")
}

func main() {
	conn, err := amqp.Dial(utils.MQURL)
	utils.FailOnError(err, utils.ErrConnectingToServer.Error())
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, utils.ErrCreatingChannel.Error())
	defer ch.Close()

	q, err := ch.QueueDeclare(`hello`, false, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringQueue.Error())

	
	body := bodyFrom(os.Args)
	err = ch.Publish(``, q.Name, false, false, amqp.Publishing{
		ContentType:  `text/plain`,
		DeliveryMode: amqp.Persistent,
		Body:         []byte(body),
	})
	utils.FailOnError(err, utils.ErrPublishingMessage.Error())
	log.Printf("[x] Sent %s", body)
}
