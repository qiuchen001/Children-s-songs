package models

import (
	"children-songs/dao"
	"github.com/jinzhu/gorm"
)

type Banner struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
	//AddTime     int64  `json:"addTime"`
	//UpdateTime  int64  `json:"updateTime"`
}

func (Banner) TableName() string {
	return "banner"
}

func GetBanners(aid int, sort string) ([]Banner, error) {
	var banners []Banner
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&banners).Error
	return banners, err
}

func GetBannerInfo(id int) (Banner, error) {
	var banner Banner
	err := dao.Db.Where("id = ?", id).First(&banner).Error
	return banner, err
}

func UpdateBanner(id int) {
	var banner Banner
	dao.Db.Model(&banner).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
