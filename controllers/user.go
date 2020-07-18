package controllers

import (
	"github.com/astaxie/beego"
	"movie/library"
	"movie/models"
	"regexp"
)

type UserController struct {
	beego.Controller
}

// @router /register/save [get]
func (uc *UserController) RegisterSave() {
	var (
		phone       string
		password    string
		err         error
		isMatched   bool
		isExisted   bool
		id          int64
	)
	phone    = uc.GetString("phone")
	password = uc.GetString("password")
	if phone == "" {
		uc.Data["json"] = ReturnError(library.ERR_NO_USER_PHONE_IS_NULL)
		uc.ServeJSON()
	}
	if isMatched, err = regexp.MatchString(`^1(3|4|5|6|7|8)[0-9]\d{8}$`, phone); !isMatched || err != nil {
		uc.Data["json"] = ReturnError(library.ERR_NO_PHONE_NOT_MATCH)
		uc.ServeJSON()
	}
	if password == "" {
		uc.Data["json"] = ReturnError(library.ERR_NO_USER_PASSWARD_IS_NULL)
		uc.ServeJSON()
	}
	if isExisted = models.IsExistedByPhone(phone); isExisted {
		uc.Data["json"] = ReturnError(library.ERR_NO_PHONE_EXISTED)
		uc.ServeJSON()
	}
	id = models.RegisterUser(phone, library.Md5(password))
	if id == -1 {
		uc.Data["json"] = ReturnError(library.ERR_NO_REGISTER_FAILED)
		uc.ServeJSON()
	}
	uc.Data["json"] = ReturnResponse(map[string]int64{"id":id})
	uc.ServeJSON()
}