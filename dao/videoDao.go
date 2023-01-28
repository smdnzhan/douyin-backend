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

// 返回VideoDao单例
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
