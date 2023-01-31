package dao

import (
	"douyin-backend/config"
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

// VideoList 根据用户id查出用户所有视频
func (*VideoDao) VideoList(userId int64) ([]entity.VideoPO, error) {
	var result []entity.VideoPO
	err := util.DB.Where(&entity.VideoPO{AuthorId: userId}).Find(&result).Error
	if err != nil {
		log.Println("查询video列表出错", err)
	}
	return result, err

}

// VideoListBefore 返回某个时间点之前的视频切片，按时间倒序排列，数量在setting.go中默认设置为5
func (*VideoDao) VideoListBefore(lastTime time.Time) ([]entity.VideoPO, error) {
	var result []entity.VideoPO
	err := util.DB.Where("publish_time<?", lastTime).
		Order("publish_time desc").
		Limit(config.FEED_COUNT).Find(&result).Error
	if err != nil {
		log.Println("查询video列表出错", err)
	}
	return result, err

}

// GetVideosByList 查询集合videoIds的所有视频
func (*VideoDao) GetVideosByList(videoIds []int64) ([]entity.VideoPO, error) {
	var result []entity.VideoPO
	err := util.DB.Model(entity.VideoPO{}).Where("id in (?)", videoIds).Find(&result).Error
	if err != nil {
		log.Println("查询video列表出错", err)
	}
	return result, err

}
