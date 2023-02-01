package dao

import (
	"douyin-backend/config"
	"douyin-backend/entity"
	"douyin-backend/util"
	"log"
	"sync"
)

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

// NewCommentDaoInstance  返回CommentDao单例
func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

// Count 使用video id 查询Comment数量
func (*CommentDao) Count(videoId int64) (int64, error) {
	log.Println("CommentDao-Count: running") //函数已运行
	var count int64
	//数据库中查询评论数量
	err := util.DB.Model(entity.Comment{}).
		Where(map[string]interface{}{"video_id": videoId, "cancel": 0}).
		Count(&count).Error
	if err != nil {
		log.Println("查询评论数量出错:", err.Error()) //函数返回提示错误信息
		return -1, err
	}
	return count, nil
}

// CommentListByVideoId 查找一个视频的所有评论
func (*CommentDao) CommentListByVideoId(videoId int64) ([]entity.Comment, error) {
	log.Println("CommentDao-Count: running") //函数已运行
	var result []entity.Comment
	err := util.DB.Model(entity.Comment{}).
		Where("video_id = ?", videoId).
		Find(&result)
	return result, err.Error
}
func (*CommentDao) InsertComment(comment entity.Comment) (entity.Comment, error) {
	err := util.DB.Model(comment).Create(&comment).Error
	if err != nil {
		log.Println("插入评论出错", err)
	}
	return comment, err
}

func (*CommentDao) DeleteComment(id int64) (entity.Comment, error) {
	var comment entity.Comment
	//确认评论是否存在
	log.Println("评论id为:", id)
	result := util.DB.Model(comment).Where("id = ?", id).First(&comment).Error
	if result != nil {
		log.Println("该评论不存在")
		return comment, result
	}
	result = util.DB.Model(comment).Where("id = ?", id).Update("cancel", config.UNCOMMENT_STATUS).Error
	return comment, result
}

// 根据评论id返回评论
func (*CommentDao) FindComment(commentId int64) (entity.Comment, error) {
	var comment entity.Comment
	err := util.DB.Model(entity.Comment{}).Where("id = ?", commentId).Find(&comment)
	return comment, err.Error
}
