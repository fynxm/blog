package controllers

import (
	"blog/models"
	"log"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	this.Preare()
	albums, err := models.FindAllAlbums()
	if err != nil {
		log.Println(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}
