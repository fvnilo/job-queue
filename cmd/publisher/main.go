package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nylo-andry/jobqueue/amqp"
)

func main() {
	hostName := os.Getenv("MQ_HOST")
	queue := os.Getenv("QUEUE_NAME")

	p := amqp.MessagePublisher{}
	p.Open(fmt.Sprintf("amqp://%v", hostName), queue)
	defer p.Close()

	for {
		if err := p.Publish([]byte("1,1")); err != nil {
			panic(err)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
