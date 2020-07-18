package models

import (
	"github.com/astaxie/beego/orm"
	"movie/library"
)

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	RegisterTime int64  `json:"register_time"`
	Status       int    `json:"status"`
	Avatar       string `json:"avatar"`
}

var (
	userOrm orm.Ormer
)

func newUserOrm() orm.Ormer {
	if userOrm == nil {
		userOrm = orm.NewOrm()
	}
	return userOrm
}

func IsExistedByPhone(phone string) bool {
	var (
		u   User
		err error
	)
	u = User{
		Phone: phone,
	}
	if err = newUserOrm().Read(&u, "Phone"); err != nil {
		return false
	}
	return true
}

func RegisterUser(phone string, password string) int64 {
	var (
		u   User
		err error
		id  int64
	)
	u = User{
		Name:         "",
		Password:     password,
		Phone:        phone,
		RegisterTime: library.NowTimeUnix(),
		Status:       library.USER_STATUS,
		Avatar:       "",
	}
	if id, err = newUserOrm().Insert(&u); err != nil {
		return -1
	}
	return id
}

func LoginUser(phone string, password string) (int64, string) {
	var (
		u      User
		err    error
	)
	if err = newUserOrm().QueryTable(library.TABLE_NAME_USER).Filter("Phone", phone).Filter("Password", password).One(&u); err != nil {
		return 0, ""
	}
	if u.Id > 0 {
		return u.Id, u.Name
	}
	return 0, ""
}

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(User))
}