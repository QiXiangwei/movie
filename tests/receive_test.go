package test

import (
	"fmt"
	"github.com/astaxie/beego"
	"movie/services/rabbitmq"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestReceive(t *testing.T) {
	rabbitmq.Costume("", "testmq", callback)
}

func callback(msg string) {
	fmt.Println(msg)
}
