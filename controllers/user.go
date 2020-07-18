package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"movie/library"
	"movie/models"
	"regexp"
)

type UserController struct {
	beego.Controller
}

// @router /user/register [get]
func (uc *UserController) UserRegister() {
	var (
		phone       string
		password    string
		id          int64
		isExisted   bool
		outputResp  *OutputResponse
	)
	phone      = uc.GetString("phone")
	password   = uc.GetString("password")
	outputResp = checkParams(phone, password)
	if isExisted = models.IsExistedByPhone(phone); isExisted {
		uc.Data["json"] = ReturnError(library.ERR_NO_PHONE_EXISTED)
		uc.ServeJSON()
	}
	if outputResp != nil {
		uc.Data["json"] = outputResp
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

// @router /user/login [get]
func (uc *UserController) UserLogin() {
	var (
		phone      string
		password   string
		outputResp *OutputResponse
		userId     int64
		name       string
	)
	phone      = uc.GetString("phone")
	password   = uc.GetString("password")

	outputResp = checkParams(phone, password)
	if outputResp != nil {
		uc.Data["json"] = outputResp
		uc.ServeJSON()
	}

	if userId, name = models.LoginUser(phone, library.Md5(password)); userId == 0 {
		uc.Data["json"] = ReturnError(library.ERR_NO_LOGIN_FAILED)
		uc.ServeJSON()
	}
	fmt.Println(name)
	fmt.Println(userId)
	uc.Data["json"] = ReturnResponse(map[string]interface{}{"uid":userId, "name":name})
	uc.ServeJSON()
}

func checkParams(phone string, password string) *OutputResponse {
	var (
		isMatched bool
		err       error
	)
	if phone == "" {
		return ReturnError(library.ERR_NO_USER_PHONE_IS_NULL)
	}
	if isMatched, err = regexp.MatchString(`^1(3|4|5|6|7|8)[0-9]\d{8}$`, phone); !isMatched || err != nil {
		return ReturnError(library.ERR_NO_PHONE_NOT_MATCH)
	}
	if password == "" {
		return ReturnError(library.ERR_NO_USER_PASSWARD_IS_NULL)
	}
	return nil
}