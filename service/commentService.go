package service

import "douyin-backend/entity"

type CommentService interface {

	// CountFromVideoId 查询视频评论数量
	CountFromVideoId(videoId int64) (int64, error)
	InsertComment(videoId int64, userId int64, commentText string) (entity.Comment, error)
	DeleteComment(userId int64, commentId int64) (entity.Comment, error)
	CommentToCommentInfo(comment *entity.Comment, info *entity.CommentInfo, userId int64)
	GetCommentById(commentId int64) (entity.Comment, error)
}
