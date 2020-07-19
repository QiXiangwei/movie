package routers

import (
	"github.com/astaxie/beego"
	"movie/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.ChannelController{})
	beego.Include(&controllers.RegionController{})
	beego.Include(&controllers.TypeController{})
}
