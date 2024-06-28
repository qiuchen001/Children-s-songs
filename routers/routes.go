package routers

import (
	"children-songs/controllers"
	"children-songs/pkg/services"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 轮播图
	user := r.Group("/banner")
	{
		bannerService := services.NewBannerService()
		bannerController := controllers.NewBannerController(bannerService)

		user.GET("/list", bannerController.List)
	}

	// 分类
	category := r.Group("/category")
	{
		category.GET("/list", controllers.NewCategoryController().List)
	}

	// 歌曲
	song := r.Group("/song")
	{
		songService := services.NewSongService()
		songController := controllers.NewSongController(songService)

		song.GET("/list", songController.List)
	}

	return r
}
