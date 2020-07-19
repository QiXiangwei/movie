package controllers

import (
	"github.com/astaxie/beego"
	"movie/library"
	"movie/models"
)

type RegionController struct {
	beego.Controller
}

// @router /region/list [get]
func (rc *RegionController) RegionList() {
	var (
		channelId  int64
		err        error
		regionList []models.Region
		total      int64
	)
	if channelId, err = rc.GetInt64("channelId"); err != nil {
		rc.Data["json"] = ReturnResponse(map[string]interface{}{"total":total, "regionList":nil})
		rc.ServeJSON()
	}
	if total, regionList = models.GetRegionByChannelId(channelId); total == 0 {
		rc.Data["json"] = ReturnResponse(map[string]interface{}{"total":total, "regionList":nil})
		rc.ServeJSON()
	}
	rc.Data["json"] = ReturnResponse(map[string]interface{}{"total":total, "regionList":regionList})
	rc.ServeJSON()
}

// @router /region/create [get]
func (rc *RegionController) RegionCreate() {
	var (
		name      string
		channelId int64
		err       error
		id        int64
		isChanel  bool
	)
	name = rc.GetString("name")
	if name == "" {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_NAME_IS_NULL)
		rc.ServeJSON()
	}
	if channelId, err = rc.GetInt64("channelId"); err != nil {
		rc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_ID_ILLEGAL)
		rc.ServeJSON()
	}
	if isChanel = models.IsExistedChannel(channelId); !isChanel {
		rc.Data["json"] = ReturnError(library.ERR_NO_CHANNEL_UNEXISTED)
		rc.ServeJSON()
	}
	if id = models.CreateRegion(name, channelId); id == 0 {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_CREATE_FAILED)
		rc.ServeJSON()
	}
	rc.Data["json"] = ReturnResponse(map[string]int64{"id":id})
	rc.ServeJSON()
}

// @router /region/online [get]
func (rc *RegionController) RegionOnline() {
	var (
		regionId  int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if regionId, err = rc.GetInt64("regionId"); err != nil {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_ID_ILLEGAL)
		rc.ServeJSON()
	}
	if isExisted = models.IsExistedRegion(regionId); !isExisted {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_UNEXISTED)
		rc.ServeJSON()
	}
	if isSucceed = models.OnlineRegion(regionId); !isSucceed {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_ONLINE_FAILED)
		rc.ServeJSON()
	}
	rc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed": isSucceed})
	rc.ServeJSON()
}

// @router /region/offline [get]
func (rc *RegionController) RegionOffline() {
	var (
		regionId  int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if regionId, err = rc.GetInt64("regionId"); err != nil {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_ID_ILLEGAL)
		rc.ServeJSON()
	}
	if isExisted = models.IsExistedRegion(regionId); !isExisted {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_UNEXISTED)
		rc.ServeJSON()
	}
	if isSucceed = models.OfflineRegion(regionId); !isSucceed {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_OFFLINE_FAILED)
		rc.ServeJSON()
	}
	rc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed": isSucceed})
	rc.ServeJSON()
}


// @router /region/delete [get]
func (rc * RegionController) RegionDelete() {
	var (
		regionId  int64
		err       error
		isExisted bool
		isSucceed bool
	)
	if regionId, err = rc.GetInt64("regionId"); err != nil {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_ID_ILLEGAL)
		rc.ServeJSON()
	}
	if isExisted = models.IsExistedRegion(regionId); !isExisted {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_UNEXISTED)
		rc.ServeJSON()
	}
	if isSucceed = models.DeleteRegion(regionId); !isSucceed {
		rc.Data["json"] = ReturnError(library.ERR_NO_REGION_DELETE_FAILED)
		rc.ServeJSON()
	}
	rc.Data["json"] = ReturnResponse(map[string]bool{"isSucceed": isSucceed})
	rc.ServeJSON()
}


