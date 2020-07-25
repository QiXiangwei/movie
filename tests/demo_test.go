package test

import (
	"fmt"
	"github.com/astaxie/beego"
	"movie/services/rabbitmq"
	"path/filepath"
	"runtime"
	"testing"
)

type TestPro struct {
	msgContent string
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func (m *TestPro) MsgContent() string {
	return m.msgContent
}

func (m *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func TestDemo(t *testing.T) {
	msg := "test"
	m := &TestPro{msg}
	queueExchange := &rabbitmq.QueueExchange{
		"test.rabbit",
		"rabbit.key",
		"test.rabbit.mq",
		"direct",
	}
	mq := rabbitmq.New(queueExchange)
	mq.NewConnect()
	mq.RegisterProducer(m)
	mq.RegisterReceiver(m)

	mq.Start()
}
