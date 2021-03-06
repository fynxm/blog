package controllers

import (
	"blog/models"
	"fmt"
)

type UpdateArticleController struct {
	BaseController
}

func (this *UpdateArticleController) Get() {
	this.Preare()
	id, _ := this.GetInt("id")
	fmt.Println("id", id)
	//获取对应文章
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"
}

//修改文章操作
func (this *UpdateArticleController) Post() {
	id, _ := this.GetInt("id")
	//获取前端传的值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	//实例化model，修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, e := models.UpdateArticle(art)
	if e != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}
	this.ServeJSON()
}
