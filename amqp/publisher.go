package amqp

import (
	"fmt"
	"log"

	"github.com/nylo-andry/jobqueue"
	"github.com/streadway/amqp"
)

var _ jobqueue.Publisher = &MessagePublisher{}

type MessagePublisher struct {
	hostname  string
	queueName string
	channel   *amqp.Channel
}

func (p *MessagePublisher) Open(hostname string, queueName string) error {
	conn, err := amqp.Dial(hostname)
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not open channel: %v", err)
		return err
	}
	p.channel = ch
	p.queueName = queueName
	return nil
}

func (p *MessagePublisher) Publish(message []byte) error {
	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         message,
	}

	if err := p.channel.Publish("", p.queueName, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] failed to publish to queue %v", err)
	}
	return nil
}

func (p *MessagePublisher) Close() {
	p.channel.Close()
}
