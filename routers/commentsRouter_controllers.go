package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["movie/controllers:ChannelController"] = append(beego.GlobalControllerRouter["movie/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChannelAll",
            Router: "/channel/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:ChannelController"] = append(beego.GlobalControllerRouter["movie/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChannelCreate",
            Router: "/channel/create",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:ChannelController"] = append(beego.GlobalControllerRouter["movie/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChannelDelete",
            Router: "/channel/delete",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:ChannelController"] = append(beego.GlobalControllerRouter["movie/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChanelOffline",
            Router: "/channel/offline",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:ChannelController"] = append(beego.GlobalControllerRouter["movie/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChannelOnline",
            Router: "/channel/online",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:RegionController"] = append(beego.GlobalControllerRouter["movie/controllers:RegionController"],
        beego.ControllerComments{
            Method: "RegionCreate",
            Router: "/region/create",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:RegionController"] = append(beego.GlobalControllerRouter["movie/controllers:RegionController"],
        beego.ControllerComments{
            Method: "RegionDelete",
            Router: "/region/delete",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:RegionController"] = append(beego.GlobalControllerRouter["movie/controllers:RegionController"],
        beego.ControllerComments{
            Method: "RegionList",
            Router: "/region/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:RegionController"] = append(beego.GlobalControllerRouter["movie/controllers:RegionController"],
        beego.ControllerComments{
            Method: "RegionOffline",
            Router: "/region/offline",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:RegionController"] = append(beego.GlobalControllerRouter["movie/controllers:RegionController"],
        beego.ControllerComments{
            Method: "RegionOnline",
            Router: "/region/online",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:TypeController"] = append(beego.GlobalControllerRouter["movie/controllers:TypeController"],
        beego.ControllerComments{
            Method: "TypeCreate",
            Router: "/type/create",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:TypeController"] = append(beego.GlobalControllerRouter["movie/controllers:TypeController"],
        beego.ControllerComments{
            Method: "TypeDelete",
            Router: "/type/delete",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:TypeController"] = append(beego.GlobalControllerRouter["movie/controllers:TypeController"],
        beego.ControllerComments{
            Method: "TypeList",
            Router: "/type/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:TypeController"] = append(beego.GlobalControllerRouter["movie/controllers:TypeController"],
        beego.ControllerComments{
            Method: "TypeOffline",
            Router: "/type/offline",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/controllers:TypeController"] = append(beego.GlobalControllerRouter["movie/controllers:TypeController"],
        beego.ControllerComments{
            Method: "TypeOnline",
            Router: "/type/online",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
