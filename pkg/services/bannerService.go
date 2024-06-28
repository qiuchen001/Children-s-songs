package services

import (
	"children-songs/models"
	"github.com/gin-gonic/gin"
)

type IBannerService interface {
	List(ctx *gin.Context) (models.Banner, error)
}

type BannerService struct {
}

func (b *BannerService) List(ctx *gin.Context) (models.Banner, error) {

	return models.Banner{}, nil
}

func NewBannerService() IBannerService {
	return &BannerService{}
}
