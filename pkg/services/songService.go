package services

import (
	"children-songs/models"
	"github.com/gin-gonic/gin"
)

type ISongService interface {
	List(ctx *gin.Context) (models.Song, error)
}

type SongService struct {
}

func (b *SongService) List(ctx *gin.Context) (models.Song, error) {

	return models.Song{}, nil
}

func NewSongService() IBannerService {
	return &SongService{}
}
