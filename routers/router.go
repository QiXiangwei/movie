package routers

import (
	"movie/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
    beego.Include(&controllers.ChannelController{})
    beego.Include(&controllers.RegionController{})
}
