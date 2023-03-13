package utils

import (
	"errors"
	"log"
)

const MQURL = `amqp://guest:guest@127.0.0.1:5672/`

var ErrConnectingToServer = errors.New(`error connecting to rabbitmq server`)
var ErrCreatingChannel = errors.New(`error creating a channel`)
var ErrDeclaringQueue = errors.New(`error declaring a queue`)
var ErrPublishingMessage = errors.New(`error publishing a message`)
var ErrRegisteringConsumer = errors.New(`error registering a consumer`)



func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
