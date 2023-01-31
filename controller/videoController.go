package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Publish 上传视频
func Publish(c *gin.Context) {
	data, err := c.FormFile("data")
	//从上下文即context中取出user_id
	user_id, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	log.Printf("获取到用户id:%v\n", user_id)
	title := c.PostForm("title")
	log.Printf("获取到视频title:%v\n", title)
	if err != nil {
		log.Println("获取视频流失败")
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	vsi := service.VideoServiceImpl{}
	//登录的
	err = vsi.Publish(data, user_id, title)
	if err != nil {
		log.Println("controller发布视频失败")
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	log.Println("controller发布视频成功")

	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList 已知bug 自己查自己会报错(follow表没有自己关注自己的)
func PublishList(c *gin.Context) {
	//目标用户Id
	targetId := c.Query("user_id")
	targetIdInt, _ := strconv.ParseInt(targetId, 10, 64)
	log.Println("目标用户id：" + targetId)
	//从上下文中获取 当前用户id
	userId := c.GetString("user_id")
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	log.Println("当前用户id：" + userId)
	usi := service.NewUserServiceImplInstance()
	var result entity.UserInfo
	if len(userId) == 0 {
		//当前用户是未登录用户
		result = usi.UNGetUserInfo(targetIdInt)
	} else {
		result = usi.GetUserInfo(userIdInt, targetIdInt)
	}
	vsi := service.NewVideoServiceImplInstance()
	csi := service.NewCommentServiceImplInstance()
	lsi := service.NewLikeServiceImplInstance()
	videoPOList, _ := vsi.PublishList(targetIdInt)
	videoInfoList := make([]entity.VideoInfo, len(videoPOList), len(videoPOList))

	//log.Println("视频基础信息列表:", videoPOList)

	//循环赋值 此处可以用多协程优化 逻辑应该封装在在videoService里..？
	for i := 0; i < len(videoPOList); i++ {
		favoriteCount, _ := lsi.FavouriteCount(videoPOList[i].Id)
		commentCount, _ := csi.CommentCountFromVideoId(videoPOList[i].Id)
		var isFavorite bool
		if userId != targetId {
			isFavorite, _ = lsi.IsFavorite(userIdInt, videoPOList[i].Id)
		} else if len(userId) == 0 {
			isFavorite = false
		}
		if userId == targetId {
			isFavorite = true
		}

		element := &entity.VideoInfo{
			VideoPO:       videoPOList[i],
			Author:        result,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
		}
		//log.Println("视频信息:", element)
		videoInfoList[i] = *element
	}
	//log.Println("视频列表:", videoInfoList)
	c.JSON(http.StatusOK, entity.VideoListResponse{
		StatusCode:    0,
		StatusMsg:     "查询用户视频列表成功",
		VideoInfoList: videoInfoList,
	})

}

func Feed(c *gin.Context) {
	//获取用户Id
	userId := c.GetString("user_id")
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	log.Println("当前用户id：" + userId)
	vsi := service.NewVideoServiceImplInstance()
	//获取时间戳
	inputTime := c.Query("latest_time")
	log.Printf("传入的时间" + inputTime)
	var lastTime time.Time
	if len(inputTime) != 0 && inputTime != "0" {
		log.Println("用户传入了时间")
		me, _ := strconv.ParseInt(inputTime, 10, 64)
		if me > time.Now().Unix() {
			lastTime = time.Now()
		} else {
			lastTime = time.Unix(me, 0)
		}

	} else {
		log.Println("用户未传入时间")
		lastTime = time.Now()
	}
	log.Printf("时间为:%v", lastTime)
	videoInfoList, err := vsi.Feed(lastTime, userIdInt)
	if err != nil {
		log.Println("feed视频流出错")
	}
	c.JSON(http.StatusOK, entity.VideoListResponse{
		StatusCode:    0,
		StatusMsg:     "查询用户视频列表成功",
		NextTime:      time.Now().Unix(),
		VideoInfoList: videoInfoList,
	})
}
