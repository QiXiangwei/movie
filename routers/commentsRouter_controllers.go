package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["movie/controllers:UserController"] = append(beego.GlobalControllerRouter["movie/controllers:UserController"],
        beego.ControllerComments{
            Method: "RegisterSave",
            Router: "/register/save",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
