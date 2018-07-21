package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nylo-andry/jobqueue/amqp"
)

func main() {
	hostName := os.Getenv("MQ_HOST")
	queue := os.Getenv("QUEUE_NAME")

	l := amqp.MessageListener{}
	l.Open(fmt.Sprintf("amqp://%v", hostName), queue)
	defer l.Close()

	msgs := l.Listen()
	log.Println("[+] Waiting for messages. To exit press CTRL+C")

	for d := range msgs {
		i1, i2 := toNums(d)
		fmt.Println(time.Now().Format("01-02-2006 15:04:05"), "::", i1+i2)

	}
}

func toNums(b []byte) (int, int) {
	s := string(b)
	ss := strings.Split(s, ",")
	i1, _ := strconv.Atoi(ss[0])
	i2, _ := strconv.Atoi(ss[1])
	return i1, i2
}
