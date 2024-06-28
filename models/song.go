package models

import (
	"children-songs/dao"
	"github.com/jinzhu/gorm"
)

type Song struct {
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

func (Song) TableName() string {
	return "song"
}

func GetSongs(aid int, sort string) ([]Song, error) {
	var songs []Song
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&songs).Error
	return songs, err
}

func GetSongInfo(id int) (Song, error) {
	var song Song
	err := dao.Db.Where("id = ?", id).First(&song).Error
	return song, err
}

func UpdateSong(id int) {
	var song Song
	dao.Db.Model(&song).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
