package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (this *ShowArticleController) Get() {
	this.Preare()
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id", id)
	//获取id对应文章
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName = "show_article.html"
}
