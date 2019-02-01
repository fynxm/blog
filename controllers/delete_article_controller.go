package controllers

import (
	"blog/models"
	"log"
)

type DeleteArticleController struct {
	BaseController
}

func (this *DeleteArticleController) Get() {
	this.Preare()
	id, _ := this.GetInt("id")
	_, e := models.DeleteArticle(id)
	if e != nil {
		log.Println(e)
	}
	this.Redirect("/", 302)
}
