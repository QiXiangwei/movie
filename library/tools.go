package library

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"hash"
	"time"
)

func Md5(str string) string {
	var (
		m hash.Hash
	)
	m = md5.New()
	m.Write([]byte(str + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(m.Sum(nil))
}

func NowTimeUnix() int64 {
	return time.Now().Unix()
}