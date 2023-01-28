package service

import (
	"douyin-backend/dao"
	"douyin-backend/util"
	uuid "github.com/satori/go.uuid"
	"log"
	"mime/multipart"
	"sync"
)

type VideoServiceImpl struct {
}

var (
	videoServiceImpl *VideoServiceImpl //controller层通过该实例变量调用service的所有业务方法。
	videoServiceOnce sync.Once         //限定该service对象为单例，节约内存。
)

// 生成并返回followServiceImpl的单例对象。
func NewVideoServiceImplInstance() *VideoServiceImpl {
	videoServiceOnce.Do(
		func() {
			videoServiceImpl = &VideoServiceImpl{}
		})
	return videoServiceImpl
}
func (videoService *VideoServiceImpl) Publish(data *multipart.FileHeader, userId int64, title string) error {
	//将视频流上传到OSS服务器，保存视频链接
	file, err := data.Open()
	defer file.Close()
	if err != nil {
		log.Printf("方法data.Open() 失败%v", err)
		return err
	}
	log.Printf("方法data.Open() 成功")
	//生成一个uuid作为视频的名字
	videoName := uuid.NewV4().String()
	log.Printf("生成视频名称%v", videoName)
	//首先存至OSS服务器
	videoUrl, err := util.UploadVideo(file, videoName)
	if err != nil {
		log.Printf("存储OSS失败%v", err)
		return err
	}
	log.Printf("存储OSS成功")
	//coverName:=videoName
	//coverUrl, err := util.UploadCover(file, coverName)
	//临时使用同一张封面:
	coverUrl := "https://zdx-shangcheng.oss-cn-hangzhou.aliyuncs.com/cover/%E5%8F%AF%E8%BE%BE%E9%B8%AD.jpg"
	//组装并持久化
	err = dao.NewVideoDaoInstance().SaveVideo(videoUrl, coverUrl, userId, title)
	if err != nil {
		log.Printf("方法dao.Save(videoName, imageName, userId) 失败%v", err)
		return err
	}
	log.Printf("方法dao.Save(videoName, imageName, userId) 成功")
	return nil
}
