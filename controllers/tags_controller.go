package controllers

import "blog/models"

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	this.Preare()
	tags := models.QueryArticleWithParam("tags")
	this.Data["Tags"] = models.HandLeTagsListData(tags)
	this.TplName = "tags.html"
}
