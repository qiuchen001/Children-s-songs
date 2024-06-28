package controllers

import (
	"github.com/gin-gonic/gin"
)

type Song struct {
	Id        uint64 `json:"id"`
	ImageUrl  string `json:"imageUrl"`
	LinkType  string `json:"linkType"`
	LinkValue string `json:"linkValue"`
}

type ISongController interface {
	List(ctx *gin.Context)
}

type SongController struct {
}

func (b SongController) List(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{})
}

func NewSongController() ISongController {
	return &SongController{}
}
