package amqp

import (
	"github.com/nylo-andry/jobqueue"
	"github.com/streadway/amqp"
)

var _ jobqueue.Listener = &MessageListener{}

type MessageListener struct {
	channel *amqp.Channel
	msgs    <-chan amqp.Delivery
}

func (l *MessageListener) Open(hostname string, queueName string) error {
	conn, err := amqp.Dial(hostname)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	c, err := ch.Consume(queue.Name, "", false, false, false, false, nil)

	l.msgs = c
	l.channel = ch
	return nil
}

func (l *MessageListener) Listen() <-chan []byte {
	messages := make(chan []byte)

	go func() {
		for d := range l.msgs {
			messages <- d.Body
			d.Ack(false)
		}
	}()

	return messages
}

func (l *MessageListener) Close() {
	l.channel.Close()
}
