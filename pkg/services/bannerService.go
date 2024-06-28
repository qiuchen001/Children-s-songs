package services

import (
	"children-songs/models"
	"github.com/gin-gonic/gin"
)

type IBannerService interface {
	List(ctx *gin.Context) (models.Song, error)
}

type BannerService struct {
}

func (b *BannerService) List(ctx *gin.Context) (models.Song, error) {

	return models.Song{}, nil
}

func NewBannerService() IBannerService {
	return &BannerService{}
}
