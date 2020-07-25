package rabbitmq

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

type CallBack func(msg string)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

func Publish(exchangeName string, queueName string, body string) {
	fmt.Println("push start" + body)
	var (
		conn    *amqp.Connection
		channel *amqp.Channel
		queue   amqp.Queue
		err     error
	)
	if conn, err = Connect(); err != nil {
		fmt.Println("mq connect failed")
		return
	}
	defer conn.Close()
	if channel, err = conn.Channel(); err != nil {
		fmt.Println("mq channel create failed")
		return
	}
	defer channel.Close()
	if queue, err = channel.QueueDeclare(queueName, false, false, false, false, nil); err != nil {
		fmt.Println("mq queue create failed")
		return
	}
	if err = channel.Publish(exchangeName, queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}); err != nil {
		fmt.Println("mq push failed")
		return
	}
	fmt.Println("push stop" + body)
	return
}

func Costume(exchangeName string, queueName string, callback CallBack) {
	var (
		conn    *amqp.Connection
		channel *amqp.Channel
		queue   amqp.Queue
		err     error
		msgList <-chan amqp.Delivery
	)
	if conn, err = Connect(); err != nil {
		fmt.Println("mq connect failed")
	}
	defer conn.Close()
	if channel, err = conn.Channel(); err != nil {
		fmt.Println("mq channel create failed")
		return
	}
	defer channel.Close()
	if queue, err = channel.QueueDeclare(queueName, false, false, false, false, nil); err != nil {
		fmt.Println("mq queue create failed")
		return
	}

	if msgList, err = channel.Consume(queue.Name, exchangeName, true, false, false, false, nil); err != nil {
		fmt.Println("get queue msg list failed")
		return
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgList {
			s := ByteToString(&(msg.Body))
			callback(*s)
		}
	}()
	fmt.Println("waiting for messages")
	<-forever
}

func ByteToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
