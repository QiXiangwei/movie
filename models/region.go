package models

import (
	"github.com/astaxie/beego/orm"
	"movie/library"
)

type Region struct {
	Id         int64  `json:"id"`
	ChannelId  int64  `json:"channel_id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Status     int    `json:"status"`
	Sort       string `json:"sort"`
}

var (
	regionOrm orm.Ormer
)

func newRegionOrm() orm.Ormer {
	if regionOrm == nil {
		regionOrm = orm.NewOrm()
	}
	return regionOrm
}

func GetRegionByChannelId(channelId int64) (int64, []Region) {
	var (
		rList []Region
		err   error
		number int64
	)
	if number, err = newRegionOrm().
		QueryTable(library.TABLE_NAME_REGION).
		Filter("ChannelId", channelId).
		All(&rList); err != nil {
		return 0, nil
	}
	return number, rList
}

func CreateRegion(name string, channelId int64) int64 {
	var (
		r   Region
		id  int64
		err error
	)
	r = Region{
		ChannelId:  channelId,
		Name:       name,
		CreateTime: library.NowTimeUnix(),
		Status:     0,
		Sort:       "",
	}
	if id, err = newRegionOrm().Insert(&r); err != nil {
		return 0
	}
	return id
}

func OnlineRegion(regionId int64) bool {
	var (
		err error
	)

	if _, err = newRegionOrm().
		QueryTable(library.TABLE_NAME_REGION).
		Filter("Id", regionId).
		Update(orm.Params{"Status": library.REGION_STATUS_ONLINE}); err != nil {
		return false
	}
	return true
}

func OfflineRegion(regionId int64) bool {
	var (
		err error
	)

	if _, err = newRegionOrm().
		QueryTable(library.TABLE_NAME_REGION).
		Filter("Id", regionId).
		Update(orm.Params{"Status": library.REGION_STATUS_OFFLINE}); err != nil {
		return false
	}
	return true
}

func DeleteRegion(regionId int64) bool {
	var(
		err error
	)

	if _, err = newRegionOrm().
		QueryTable(library.TABLE_NAME_REGION).
		Filter("Id", regionId).
		Delete(); err != nil {
		return false
	}
	return true
}

func IsExistedRegion(regionId int64) bool {
	return newRegionOrm().QueryTable(library.TABLE_NAME_REGION).Filter("Id", regionId).Exist()
}

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Region))
}
