package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 27.01上传功能还需要用户登录状态校验
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
