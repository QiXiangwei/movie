package controllers

import (
	"github.com/astaxie/beego"
	"movie/library"
	"movie/models"
)

type TypeController struct {
	beego.Controller
}

// @router /type/list [get]
func (tc *TypeController) TypeList() {
	var (
		channelId int64
		err       error
		typeList  []models.Type
		total     int64
	)
	if channelId, err = tc.GetInt64("channelId"); err != nil {
		tc.Data["json"] = ReturnResponse(map[string]interface{}{"total": total, "typeList": nil})
		tc.ServeJSON()
	}
	if total, typeList = models.GetTypeByChannelId(channelId); total == 0 {
		tc.Data["json"] = ReturnResponse(map[string]interface{}{"total": total, "typeList": nil})
		tc.ServeJSON()
	}
	tc.Data["json"] = ReturnResponse(map[string]interface{}{"total": total, "typeList": typeList})
	tc.ServeJSON()
}

// @router /type/create [get]
func (tc *TypeController) TypeCreate() {
	var (
		name      string
		id        int64
		err       error
		channelId int64
		isChanel  bool
	)
	name = tc.GetString("name")
	if name == "" {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_NAME_IS_NULL)
		tc.ServeJSON()
	}
	if channelId, err = tc.GetInt64("channelId"); err != nil {
		tc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_ID_ILLEGAL)
		tc.ServeJSON()
	}
	if isChanel = models.IsExistedChannel(channelId); !isChanel {
		tc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_UNEXISTED)
		tc.ServeJSON()
	}
	if id = models.CreateType(name, channelId); id == 0 {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_CREATE_FAILED)
		tc.ServeJSON()
	}
	tc.Data["json"] = ReturnResponse(map[string]int64{"id": id})
	tc.ServeJSON()
}

// @router /type/online [get]
func (tc *TypeController) TypeOnline() {
	var (
		typeId    int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if typeId, err = tc.GetInt64("typeId"); err != nil {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_ID_ILLEGAL)
		tc.ServeJSON()
	}
	if isExisted = models.IsExistedType(typeId); !isExisted {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_UNEXISTED)
		tc.ServeJSON()
	}
	if isSucceed = models.OnlineType(typeId); !isSucceed {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_ONLINE_FAILED)
		tc.ServeJSON()
	}
	tc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed": isSucceed})
	tc.ServeJSON()
}

// @router /type/offline [get]
func (tc *TypeController) TypeOffline() {
	var (
		typeId    int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if typeId, err = tc.GetInt64("typeId"); err != nil {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_ID_ILLEGAL)
		tc.ServeJSON()
	}
	if isExisted = models.IsExistedType(typeId); !isExisted {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_UNEXISTED)
		tc.ServeJSON()
	}
	if isSucceed = models.OfflineType(typeId); !isSucceed {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_OFFLINE_FAILED)
		tc.ServeJSON()
	}
	tc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed": isSucceed})
	tc.ServeJSON()
}

// @router /type/delete [get]
func (tc *TypeController) TypeDelete() {
	var (
		typeId    int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if typeId, err = tc.GetInt64("typeId"); err != nil {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_ID_ILLEGAL)
		tc.ServeJSON()
	}
	if isExisted = models.IsExistedType(typeId); !isExisted {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_UNEXISTED)
		tc.ServeJSON()
	}
	if isSucceed = models.DeleteType(typeId); !isSucceed {
		tc.Data["json"] = ReturnError(library.ERR_NO_TYPE_DELETE_FAILED)
		tc.ServeJSON()
	}
	tc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed": isSucceed})
	tc.ServeJSON()
}
