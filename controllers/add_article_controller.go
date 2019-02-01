package controllers

import (
	"blog/models"
	"fmt"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get() {
	this.Preare()
	this.TplName = "write_article.html"
}

//添加文章

func (this *AddArticleController) Post() {
	//获取数据
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)
	//向数据库添加数据
	_, e := models.InsertArticle(models.Article{0, title, tags, short, content, "孔壹学院", time.Now().Unix()})
	//返回数据给浏览器
	var response map[string]interface{}
	if e == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()
}
