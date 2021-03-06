package models

import (
	"github.com/astaxie/beego/orm"
	"movie/library"
)

type Channel struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Status     int    `json:"status"`
}

var (
	channelOrm orm.Ormer
)

func newChannelOrm() orm.Ormer {
	if channelOrm ==  nil {
		channelOrm = orm.NewOrm()
	}
	return  channelOrm
}

func AllChannel() (int64, []Channel) {
	var (
		cList  []Channel
		number int64
		err    error
	)
	if number, err = newChannelOrm().QueryTable(library.TABLE_NAME_CHANNEL).All(&cList); err != nil {
		return 0, cList
	}
	return number, cList
}

func CreateChannel(name string) int64 {
	var (
		c   Channel
		id  int64
		err error
	)
	c = Channel{
		Name:       name,
		CreateTime: library.NowTimeUnix(),
		Status:     library.CHANNEL_STATUS_OFFLINE,
	}
	if id, err = newChannelOrm().Insert(&c); err != nil {
		id = 0
	}
	return id
}

func OnlineChannel(channelId int64) bool {
	var (
		err error
	)
	if _, err = newChannelOrm().
		QueryTable(library.TABLE_NAME_CHANNEL).
		Filter("Id", channelId).
		Update(orm.Params{"Status": library.CHANNEL_STATUS_ONLINE}); err != nil {
		return false
	}
	return true
}

func OfflineChannel(channelId int64) bool {
	var (
		err error
	)
	if _, err = newChannelOrm().
		QueryTable(library.TABLE_NAME_CHANNEL).
		Filter("Id", channelId).
		Update(orm.Params{"Status": library.CHANNEL_STATUS_OFFLINE}); err != nil {
		return false
	}
	return true
}

func DeleteChannel(channelId int64) bool {
	var (
		err error
	)
	if _, err = newChannelOrm().
		QueryTable(library.TABLE_NAME_CHANNEL).
		Filter("Id", channelId).
		Delete(); err != nil {
		return false
	}
	return true
}

func IsExistedChannel(channelId int64) bool {
	return newChannelOrm().QueryTable(library.TABLE_NAME_CHANNEL).Filter("Id", channelId).Exist()
}

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Channel))
}