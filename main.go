package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "movie/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		err error
	)
	defaultDb := beego.AppConfig.String("defaultdb")
	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		fmt.Println("*")
		fmt.Println(err.Error())
	}
	if err = orm.RegisterDataBase("default", "mysql", defaultDb); err != nil {
		fmt.Println("**")
		fmt.Println(err.Error())
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

