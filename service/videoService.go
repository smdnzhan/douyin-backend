package service

import "mime/multipart"

type VideoService interface {
	// Publish 上传封面/图片至oss服务器，同时将记录存储至mysql
	Publish(data *multipart.FileHeader, userId int64, title string) error
}
