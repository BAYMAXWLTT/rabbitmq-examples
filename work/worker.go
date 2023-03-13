package main

import (
	"bytes"
	"log"
	"time"
	"work/utils"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(utils.MQURL)
	utils.FailOnError(err, utils.ErrConnectingToServer.Error())

	ch, err := conn.Channel()
	utils.FailOnError(err, utils.ErrCreatingChannel.Error())

	// fair dispatch, doesn't dispatch a msg to a worker that hasn't acknowledged the previous message
	err = ch.Qos(1, 0, false)
	utils.FailOnError(err, `error setting Qos`)

	q, err := ch.QueueDeclare(`hello`, false, false, false, false, nil)
	utils.FailOnError(err, utils.ErrDeclaringQueue.Error())

	msgs, err := ch.Consume(q.Name, ``, true, false, false, false, nil)
	utils.FailOnError(err, utils.ErrRegisteringConsumer.Error())

	forever := make(chan interface{})
	go func() {
		for m := range msgs {
			dotCnt := bytes.Count(m.Body, []byte(`.`))
			log.Printf("Received a message: %s that takes %d seconds to complete", string(m.Body), dotCnt)
			time.Sleep(time.Duration(dotCnt) * time.Second)
			log.Printf(`done\n`)
		}
	}()
	log.Println(`[*] waiting for messages, to exit press CTRL+C`)
	<-forever
}
