package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Follow(c *gin.Context) {
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	actionType := c.Query("action_type")
	targetIdStr := c.Query("to_user_id")
	targetId, _ := strconv.ParseInt(targetIdStr, 10, 64)
	var err error
	if actionType == "1" {
		err = service.NewFollowServiceImpInstance().UpdateFollow(userId, targetId)
	} else {
		err = service.NewFollowServiceImpInstance().UpdateFollow(userId, targetId)
	}
	if err != nil {
		log.Printf("关注出错", err)
	}
	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "Success",
	})
}

func FollowList(c *gin.Context) {
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	targetIdStr := c.Query("user_id")
	targetId, _ := strconv.ParseInt(targetIdStr, 10, 64)
	userInfoList, err := service.NewFollowServiceImpInstance().FollowList(userId, targetId)
	if err != nil {
		log.Println("获取关注列表出错:", err)
	}
	c.JSON(http.StatusOK, entity.UserInfoListResponse{
		StatusCode:   0,
		StatusMsg:    "Success",
		UserInfoList: userInfoList,
	})
}
