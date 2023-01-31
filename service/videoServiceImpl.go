package service

import (
	"douyin-backend/config"
	"douyin-backend/dao"
	"douyin-backend/entity"
	"douyin-backend/util"
	uuid "github.com/satori/go.uuid"
	"log"
	"mime/multipart"
	"sync"
	"time"
)

type VideoServiceImpl struct {
}

var (
	videoServiceImpl *VideoServiceImpl //controller层通过该实例变量调用service的所有业务方法。
	videoServiceOnce sync.Once         //限定该service对象为单例，节约内存。
)

// NewVideoServiceImplInstance 生成并返回followServiceImpl的单例对象。
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

func (videoService *VideoServiceImpl) PublishList(userId int64) ([]entity.VideoPO, error) {
	resultList, err := dao.NewVideoDaoInstance().VideoList(userId)
	if err != nil {
		log.Println("查询视频列表出错")
	}
	return resultList, err
}

// Feed 返回视频feed流，userId是当前登录的用户id
func (videoService *VideoServiceImpl) Feed(lastTime time.Time, userId int64) ([]entity.VideoInfo, error) {
	videoList, err := dao.NewVideoDaoInstance().VideoListBefore(lastTime)
	len := len(videoList)
	videoInfoList := make([]entity.VideoInfo, len, len)
	for index, _ := range videoList {
		videoService.VideoPOToVideoInfo(&videoList[index], &videoInfoList[index], userId)
	}
	return videoInfoList, err
}

// VideoPOToVideoInfo 由VideoPO组装VideoInfo
func (videoService *VideoServiceImpl) VideoPOToVideoInfo(video *entity.VideoPO, videoInfo *entity.VideoInfo,
	userId int64) {
	var wg sync.WaitGroup
	wg.Add(4)
	var err error
	videoInfo.VideoPO = *video
	usi := NewUserServiceImplInstance()
	lsi := NewLikeServiceImplInstance()
	csi := NewCommentServiceImplInstance()
	//插入UserInfo 暂时不设置报错
	go func() {
		if userId != config.UNLOGIN_USER {
			videoInfo.Author = usi.GetUserInfo(userId, videoInfo.VideoPO.AuthorId)
		} else {
			videoInfo.Author = usi.UNGetUserInfo(videoInfo.VideoPO.AuthorId)
		}
		wg.Done()
	}()
	//查询点赞数量
	go func() {
		videoInfo.FavoriteCount, err = lsi.FavouriteCount(videoInfo.VideoPO.Id)
		log.Printf("视频点赞数量为:", videoInfo.FavoriteCount)
		if err != nil {
			log.Println("组装VideoInfo:查询点赞数量错误")
		}
		wg.Done()
	}()
	//查询当前用户是否点赞了该视频
	go func() {
		if userId != config.UNLOGIN_USER {
			videoInfo.IsFavorite, err = lsi.IsFavorite(userId, videoInfo.VideoPO.Id)
		} else {
			videoInfo.IsFavorite = false
		}
		wg.Done()
	}()
	//查询评论数量
	go func() {
		videoInfo.CommentCount, err = csi.CommentCountFromVideoId(videoInfo.VideoPO.Id)
		wg.Done()
	}()
	log.Printf("组装VideoInfo完成")
	wg.Wait()
}

func (videoService *VideoServiceImpl) GetVideoListByIds(videoIds []int64) ([]entity.VideoPO, error) {
	return dao.NewVideoDaoInstance().GetVideosByList(videoIds)
}
