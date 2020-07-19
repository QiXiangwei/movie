package models

import (
	"github.com/astaxie/beego/orm"
	"movie/library"
)

type Type struct {
	Id         int64  `json:"id"`
	ChannelId  int64  `json:"channel_id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Status     int    `json:"status"`
	Sort       string `json:"sort"`
}

var (
	typeOrm orm.Ormer
)

func newTypeOrm() orm.Ormer {
	if typeOrm == nil {
		typeOrm = orm.NewOrm()
	}
	return typeOrm
}

func GetTypeByChannelId(channelId int64) (int64, []Type) {
	var (
		number int64
		tList  []Type
		err    error
	)
	if number, err = newTypeOrm().
		QueryTable(library.TABLE_NAME_TYPE).
		Filter("ChannelId", channelId).
		All(&tList); err != nil {
		return 0, nil
	}
	return number, tList
}

func CreateType(name string, channelId int64) int64 {
	var (
		id  int64
		err error
		t   Type
	)
	t = Type{
		ChannelId:  channelId,
		Name:       name,
		CreateTime: library.NowTimeUnix(),
		Status:     0,
		Sort:       "",
	}
	if id, err = newTypeOrm().Insert(&t); err != nil {
		return 0
	}
	return id
}

func OnlineType(typeId int64) bool {
	var (
		err error
	)
	if _, err = newTypeOrm().
		QueryTable(library.TABLE_NAME_TYPE).
		Filter("Id", typeId).
		Update(orm.Params{"Status": library.TYPE_STATUS_ONLINE}); err != nil {
		return false
	}
	return true
}

func OfflineType(typeId int64) bool {
	var (
		err error
	)
	if _, err = newTypeOrm().
		QueryTable(library.TABLE_NAME_TYPE).
		Filter("Id", typeId).
		Update(orm.Params{"Status": library.TYPE_STATUS_OFFLINE}); err != nil {
		return false
	}
	return true
}

func DeleteType(typeId int64) bool {
	var (
		err error
	)
	if _, err = newTypeOrm().
		QueryTable(library.TABLE_NAME_TYPE).
		Filter("Id", typeId).
		Delete(); err != nil {
		return false
	}
	return true
}

func IsExistedType(typeId int64) bool {
	return newTypeOrm().QueryTable(library.TABLE_NAME_TYPE).Filter("Id", typeId).Exist()
}

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Type))
}
