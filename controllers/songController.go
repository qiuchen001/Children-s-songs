package controllers

import (
	"children-songs/pkg/services"
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
	SongService services.ISongService
}

func (b SongController) List(ctx *gin.Context) {
	banner, err := b.SongService.List(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, banner)
}

func NewSongController(s services.ISongService) ISongController {
	return &SongController{SongService: s}
}
