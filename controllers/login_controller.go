package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	userName := this.GetString("username")
	passWord := this.GetString("password")

	fmt.Println("username:", userName, "password", passWord)

	//MD5转换密码
	newPass := utils.MD5(passWord)
	id := models.QueryUserWithParam(userName, newPass)
	if id > 0 {
		this.SetSession("loginuser", userName)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登陆成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
	return
}
