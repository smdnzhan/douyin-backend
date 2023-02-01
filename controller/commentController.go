package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	//从上下文中去userId
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentIdStr := c.Query("comment_id")
	commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
	var comment entity.Comment
	var err error
	if actionType == "1" {
		comment, err = service.NewCommentServiceImplInstance().InsertComment(videoId, userId, commentText)
		if err != nil {
			log.Printf("新增评论出错")
		}
	} else {
		comment, err = service.NewCommentServiceImplInstance().DeleteComment(userId, commentId)
		if err != nil {
			log.Printf("删除评论出错")
		}
	}
	var commentInfo entity.CommentInfo
	service.NewCommentServiceImplInstance().CommentToCommentInfo(&comment, &commentInfo, userId)
	log.Println("commentInfo组装结果:", commentInfo)
	c.JSON(http.StatusOK, entity.CommentInfoResponse{
		StatusCode:  0,
		StatusMsg:   "Success",
		CommentInfo: commentInfo,
	})
}

func CommentList(c *gin.Context) {
	//从上下文中去找userId
	userIdStr := c.GetString("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentInfoList, err := service.NewCommentServiceImplInstance().GetCommentsByVideoId(videoId, userId)
	if err != nil {
		log.Printf("查询评论列表出错")
	}
	c.JSON(http.StatusOK, entity.CommentInfoListResponse{
		StatusCode:      0,
		StatusMsg:       "Success",
		CommentInfoList: commentInfoList,
	})
}
