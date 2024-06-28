package controllers

import (
	"github.com/gin-gonic/gin"
)

type Category struct {
	Id        uint64 `json:"id"`
	ImageUrl  string `json:"imageUrl"`
	LinkType  string `json:"linkType"`
	LinkValue string `json:"linkValue"`
}

type ICategoryController interface {
	List(ctx *gin.Context)
}

type CategoryController struct {
}

func (b CategoryController) List(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{})
}

func NewCategoryController() ICategoryController {
	return &CategoryController{}
}
