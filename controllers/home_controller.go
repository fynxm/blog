package controllers

import (
	"blog/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Preare()
	tag := this.GetString("tag")
	fmt.Println("tag:", tag)
	page, _ := this.GetInt("page")
	var artList []models.Article

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		this.Data["HasFooter"] = false

	} else {
		if page <= 0 {
			page = 1
		}
		artList, _ = models.FindArticleWithpage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}
