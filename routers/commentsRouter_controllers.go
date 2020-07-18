package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["movie/controllers:UserController"] = append(beego.GlobalControllerRouter["movie/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserLogin",
            Router: "/user/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:UserController"] = append(beego.GlobalControllerRouter["movie/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserRegister",
            Router: "/user/register",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
