package utils

import (
	"errors"
	"log"
	"math/rand"
)

const MQURL = `amqp://guest:guest@127.0.0.1:5672/`

var ErrConnectingToServer = errors.New(`error connecting to rabbitmq server`)
var ErrCreatingChannel = errors.New(`error creating a channel`)
var ErrDeclaringQueue = errors.New(`error declaring a queue`)
var ErrDeclaringExchange = errors.New(`err declaring an exchange`)
var ErrPublishingMessage = errors.New(`error publishing a message`)
var ErrRegisteringConsumer = errors.New(`error registering a consumer`)
var ErrBindingQueue = errors.New(`error binding queue`)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}

func Fibo(n int) int {
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i < n+1; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n]
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
