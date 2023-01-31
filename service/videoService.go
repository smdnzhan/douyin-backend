package service

import (
	"douyin-backend/entity"
	"mime/multipart"
)

type VideoService interface {
	// Publish 上传封面/图片至oss服务器，同时将记录存储至mysql
	Publish(data *multipart.FileHeader, userId int64, title string) error

	PublishList(userId int64) ([]entity.VideoPO, error)

	VideoPOToVideoInfo(video *entity.VideoPO, videoInfo *entity.VideoInfo, userId int64)

	GetVideoListByIds(videoIds []int64) ([]entity.VideoPO, error)
}
