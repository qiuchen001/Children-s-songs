package controllers

import (
	"children-songs/pkg/services"
	"github.com/gin-gonic/gin"
)

type IBannerController interface {
	List(ctx *gin.Context)
}

type BannerController struct {
	BannerService services.IBannerService
}

func (b *BannerController) List(ctx *gin.Context) {
	banner, err := b.BannerService.List(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, banner)
}

func NewBannerController(s services.IBannerService) IBannerController {
	return &BannerController{BannerService: s}
}
