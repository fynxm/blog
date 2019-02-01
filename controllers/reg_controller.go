package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {
	this.TplName = "register.html"
}

//注册处理逻辑

func (this *RegController) Post() {
	//获取表单属性
	userName := this.GetString("username")
	passWord := this.GetString("password")
	fmt.Println(userName, passWord)

	//查询数据库是否存在用户名
	id := models.QueryUserWithUsername(userName)

	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已存在"}
		this.ServeJSON()
		return
	}
	//不存在将数据保存
	newPassWord := utils.MD5(passWord)

	user := models.User{0, userName, newPassWord, 0, time.Now().Unix()}
	//保存数据
	_, e := models.InsertUser(user)
	if e != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	this.ServeJSON()
	return
}
