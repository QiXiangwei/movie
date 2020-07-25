package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var (
	mqConn *amqp.Connection
	mqChan *amqp.Channel
)

type Producer interface {
	MsgContent() string
}

type Receiver interface {
	Consumer([]byte) error
}

type RabbitMq struct {
	mqConnection *amqp.Connection
	mqChannel    *amqp.Channel
	QueueName    string
	RoutingKey   string
	ExchangeName string
	ExchangeType string
	producerList []Producer
	receiverList []Receiver
	mu           sync.RWMutex
}

type QueueExchange struct {
	QuName string
	RtKey  string
	ExName string
	ExType string
}

func (r *RabbitMq) NewConnect() {
	r.mqConnect()
}

func (r *RabbitMq) mqConnect() {
	var (
		err error
	)
	if mqConn, err = amqp.Dial("amqp://guest:guest@127.0.0.1:5672/"); err != nil {
		fmt.Println(err.Error() + "connection failed")
		return
	}
	if mqChan, err = mqConn.Channel(); err != nil {
		fmt.Println(err.Error() + "create channel failed")
		return
	}
	r.mqChannel = mqChan
}

func (r *RabbitMq) mqClose() {
	var (
		err error
	)
	if err = r.mqChannel.Close(); err != nil {
		fmt.Println(err.Error() + "channel close failed")
		return
	}
	if err = r.mqConnection.Close(); err != nil {
		fmt.Println(err.Error() + "connection close failed")
		return
	}
	return
}

func New(q *QueueExchange) *RabbitMq {
	return &RabbitMq{
		QueueName:    q.QuName,
		RoutingKey:   q.RtKey,
		ExchangeName: q.ExName,
		ExchangeType: q.ExType,
	}
}

func (r *RabbitMq) Start() {
	for _, producer := range r.producerList {
		go r.listenProducer(producer)
	}
	for _, receiver := range r.receiverList {
		go r.listenReceiver(receiver)
	}
	time.Sleep(5 * time.Second)
}

func (r *RabbitMq) RegisterProducer(producer Producer) {
	r.producerList = append(r.producerList, producer)
}

func (r *RabbitMq) listenProducer(producer Producer) {
	var (
		err error
	)
	if r.mqChannel == nil {
		r.mqConnect()
	}
	if _, err = r.mqChannel.QueueDeclarePassive(r.QueueName, true, false, false, true, nil); err != nil {
		fmt.Println("queue is not existed" + r.QueueName)
		if _, err = r.mqChannel.QueueDeclare(r.QueueName, true, false, false, true, nil); err != nil {
			fmt.Println("queue create failed" + r.QueueName)
			fmt.Println(err.Error())
			return
		}
	}

	if err = r.mqChannel.QueueBind(r.QueueName, r.RoutingKey, r.ExchangeName, true, nil); err != nil {
		fmt.Println("queue bind failed" + r.QueueName)
		fmt.Println(err.Error())
		return
	}

	if err = r.mqChannel.ExchangeDeclarePassive(r.ExchangeName, r.ExchangeType, true, false, false, true, nil); err != nil {
		fmt.Println("exchange is not existed" + r.ExchangeName)
		if err = r.mqChannel.ExchangeDeclare(r.ExchangeName, r.ExchangeType, true, false, false, true, nil); err != nil {
			fmt.Println("exchange create failed" + r.ExchangeName)
			fmt.Println(err.Error())
			return
		}
	}

	if err = r.mqChannel.Publish(r.ExchangeName, r.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(producer.MsgContent()),
	}); err != nil {
		fmt.Println("producer push failed")
		fmt.Println(err.Error())
		return
	}
}

func (r *RabbitMq) RegisterReceiver(receiver Receiver) {
	r.mu.Lock()
	r.receiverList = append(r.receiverList, receiver)
	r.mu.Unlock()
}

func (r *RabbitMq) listenReceiver(receiver Receiver) {
	var (
		err error
	)
	defer r.mqClose()
	if r.mqChannel == nil {
		r.mqConnect()
	}

	if _, err = r.mqChannel.QueueDeclarePassive(r.QueueName, true, false, false, true, nil); err != nil {
		fmt.Println("r queue is not existed" + r.QueueName)
		if _, err = r.mqChannel.QueueDeclare(r.QueueName, true, false, false, true, nil); err != nil {
			fmt.Println("r queue create failed: " + r.QueueName)
			fmt.Println(err.Error())
			return
		}
	}

	if err = r.mqChannel.QueueBind(r.QueueName, r.RoutingKey, r.ExchangeName, true, nil); err != nil {
		fmt.Println("r queue bind failed" + r.QueueName)
		fmt.Println(err.Error())
		return
	}

	if err = r.mqChannel.Qos(1, 0, true); err != nil {
		fmt.Println(err.Error())
		return
	}
	msgList, err := r.mqChannel.Consume(r.QueueName, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for msg := range msgList {
		err = receiver.Consumer(msg.Body)
		if err != nil {
			if err = msg.Ack(true); err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			if err = msg.Ack(false); err != nil {
				fmt.Println(err.Error())
				return
			}
			return
		}
	}
}
