package dao

import (
	"douyin-backend/entity"
	"douyin-backend/util"
	"log"
	"sync"
	"time"
)

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

// NewVideoDaoInstance 返回VideoDao单例
func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) SaveVideo(playUrl string, coverUrl string, authorId int64, title string) error {
	var video entity.VideoPO
	video.PublishTime = time.Now()
	video.PlayUrl = playUrl
	video.CoverUrl = coverUrl
	video.AuthorId = authorId
	video.Title = title
	err := util.DB.Save(&video).Error
	if err != nil {
		log.Printf("存储视频出错" + err.Error())
	}
	return err
}

func (*VideoDao) VideoList(userId int64) ([]entity.VideoPO, error) {
	var result []entity.VideoPO
	err := util.DB.Where(&entity.VideoPO{AuthorId: userId}).Find(&result).Error
	if err != nil {
		log.Println("查询video列表出错", err)
	}
	return result, err

}
