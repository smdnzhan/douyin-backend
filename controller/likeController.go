package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 点赞/取消点赞操作
func Favorite(c *gin.Context) {
	//从上下文中去userId
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoIdStr := c.Query("video_id")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	lsi := service.NewLikeServiceImplInstance()
	err := lsi.UpdateLike(videoId, userId, actionType)
	if err != nil {
		log.Printf("点赞出现错误")
		c.Abort()
		c.JSON(http.StatusUnauthorized, entity.Response{
			StatusCode: -1,
			StatusMsg:  "点赞出现错误",
		})
	}

	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "点赞用户视频成功",
	})

}

func FavoriteList(c *gin.Context) {
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	lsi := service.NewLikeServiceImplInstance()
	videoInfoList, err := lsi.GetLikedVideoList(userId)
	if err != nil {
		c.JSON(http.StatusNoContent, entity.VideoListResponse{
			StatusCode:    -1,
			StatusMsg:     "查询失败",
			VideoInfoList: videoInfoList,
		})
	} else {
		c.JSON(http.StatusOK, entity.VideoListResponse{
			StatusCode:    0,
			StatusMsg:     "查询点赞列表成功",
			VideoInfoList: videoInfoList,
		})
	}
}
