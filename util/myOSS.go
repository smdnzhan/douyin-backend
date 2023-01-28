package util

import (
	"douyin-backend/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
)

var OSSClient *oss.Client

func handleError(err error) {
	fmt.Println("OSSError:", err)
}
func InitOSS() {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	endpoint := config.OSS_ENDPOINT
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := config.OSS_ACCESSKEYID
	accessKeySecret := config.OSS_ACCESSSECRET
	var err error
	// 创建OSSClient实例。
	OSSClient, err = oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}

}

// 文件URL的格式为https://BucketName.Endpoint/ObjectName
// ObjectName需填写包含文件夹以及文件后缀在内的该文件的完整路径
func UploadVideo(file io.Reader, videoName string) (string, error) {
	bucket, err := OSSClient.Bucket(config.OSS_BUCKET)
	if err != nil {
		handleError(err)
	}
	//上传文件流
	object_name := config.VIDEO_PATH + "/" + videoName + ".mp4"
	err = bucket.PutObject(object_name, file)
	videoUrl := "https://" + config.OSS_BUCKET + "." + config.OSS_ENDPOINT + "/" + object_name
	if err != nil {
		handleError(err)
	}
	log.Println("视频存入成功:" + videoUrl)
	return videoUrl, err
}

func UploadCover(file io.Reader, coverName string) (string, error) {
	bucket, err := OSSClient.Bucket(config.OSS_BUCKET)
	if err != nil {
		handleError(err)
	}
	object_name := config.COVER_PATH + "/" + coverName
	//上传文件流
	err = bucket.PutObject(object_name, file)

	if err != nil {
		handleError(err)
	}
	coverUrl := "https://" + config.OSS_BUCKET + "." + config.OSS_ENDPOINT + "/" + object_name
	if err != nil {
		handleError(err)
	}
	log.Println("封面存入成功:" + coverUrl)
	return coverUrl, err
}
