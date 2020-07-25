package test

import (
	"github.com/astaxie/beego"
	"movie/services/rabbitmq"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestPush(t *testing.T) {
	count := 0

	for i := 0; i < 10; i++ {
		rabbitmq.Publish("", "testmq", "test"+strconv.Itoa(count))
		count++
		time.Sleep(5 * time.Second)
	}
}
