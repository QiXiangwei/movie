package controllers

import (
	"github.com/astaxie/beego"
	"movie/library"
	"movie/models"
)

type ChannelController struct {
	beego.Controller
}

// @router /channel/all [get]
func (cc *ChannelController) ChannelAll() {
	var (
		number      int64
		channelList []models.Channel
	)
	if number, channelList = models.AllChannel(); number == 0 {
		cc.Data["json"] = ReturnResponse(map[string]interface{}{"total":number, "channelList":nil})
		cc.ServeJSON()
	}
	cc.Data["json"] = ReturnResponse(map[string]interface{}{"total":number, "channelList":channelList})
	cc.ServeJSON()
}

// @router /channel/create [get]
func (cc *ChannelController) ChannelCreate() {
	var (
		name string
		id   int64
	)
	name  = cc.GetString("name")
	if name == "" {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_NAME_IS_NULL)
		cc.ServeJSON()
	}
	if id = models.CreateChannel(name); id == 0 {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_CREATE_FAILED)
		cc.ServeJSON()
	}
	cc.Data["json"] = ReturnResponse(map[string]int64{"id":id})
	cc.ServeJSON()
}

// @router /channel/online [get]
func (cc *ChannelController) ChannelOnline() {
	var(
		id        int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if id, err = cc.GetInt64("channelId"); err != nil {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_ID_ILLEGAL)
		cc.ServeJSON()
	}
	if isExisted = models.IsExistedChannel(id); !isExisted {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_UNEXISTED)
		cc.ServeJSON()
	}
	if isSucceed = models.OnlineChannel(id); !isSucceed {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_ONLINE_FAILED)
		cc.ServeJSON()
	}
	cc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed":isSucceed})
	cc.ServeJSON()
}

// @router /channel/offline [get]
func (cc *ChannelController) ChanelOffline() {
	var(
		id        int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if id, err = cc.GetInt64("channelId"); err != nil {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_ID_ILLEGAL)
		cc.ServeJSON()
	}
	if isExisted = models.IsExistedChannel(id); !isExisted {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_UNEXISTED)
		cc.ServeJSON()
	}
	if isSucceed = models.OfflineChannel(id); !isSucceed {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_OFFLINE_FAILED)
		cc.ServeJSON()
	}
	cc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed":isSucceed})
	cc.ServeJSON()
}

// @router /channel/delete [get]
func (cc *ChannelController) ChannelDelete() {
	var (
		id        int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if id, err = cc.GetInt64("channelId"); err != nil {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_ID_ILLEGAL)
		cc.ServeJSON()
	}
	if isExisted = models.IsExistedChannel(id); !isExisted {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_UNEXISTED)
		cc.ServeJSON()
	}
	if isSucceed = models.DeleteChannel(id); !isSucceed {
		cc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_DELETE_FAILED)
		cc.ServeJSON()
	}
	cc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed":isSucceed})
	cc.ServeJSON()
}
