package dao

import (
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

// Count 1、使用video id 查询Comment数量
func (*CommentDao) Count(videoId int64) (int64, error) {
	log.Println("CommentDao-Count: running") //函数已运行
	//Init()
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
