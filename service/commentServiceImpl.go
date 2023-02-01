package service

import (
	"douyin-backend/dao"
	"douyin-backend/entity"
	"log"
	"sync"
	"time"
)

type CommentServiceImpl struct {
	UserService
}

var (
	commentServiceImpl *CommentServiceImpl //controller层通过该实例变量调用service的所有业务方法。
	commentServiceOnce sync.Once           //限定该service对象为单例，节约内存。
)

// NewCommentServiceImplInstance 生成并返回CommentServiceImpl的单例对象。
func NewCommentServiceImplInstance() *CommentServiceImpl {
	commentServiceOnce.Do(
		func() {
			commentServiceImpl = &CommentServiceImpl{}
		})
	return commentServiceImpl
}

// 查询评论数量
func (c CommentServiceImpl) CommentCountFromVideoId(videoId int64) (int64, error) {
	//先在缓存中查
	//2.缓存中查不到则去数据库查
	cntDao, err1 := dao.NewCommentDaoInstance().Count(videoId)
	log.Println("视频评论数量 :", cntDao)
	if err1 != nil {
		log.Println("查询评论数量错误:", err1)
		return 0, nil
	}
	//返回结果
	return cntDao, nil
}

func (c CommentServiceImpl) InsertComment(videoId int64, userId int64, commentText string) (entity.Comment, error) {
	comment := entity.Comment{
		VideoId:     videoId,
		UserId:      userId,
		CommentText: commentText,
		CreateDate:  time.Now(),
		Cancel:      0,
	}
	result, err := dao.NewCommentDaoInstance().InsertComment(comment)
	if err != nil {
		log.Println("新增评论出错")
	}
	return result, err
}

func (c CommentServiceImpl) DeleteComment(userId int64, commentId int64) (entity.Comment, error) {
	var result entity.Comment
	comment, err := dao.NewCommentDaoInstance().FindComment(commentId)
	log.Println("目标评论id:", commentId)
	if err != nil {
		log.Println("查找评论出错")
	} else {
		if comment.UserId != userId {
			log.Println("当前用户不是评论者，不能删除")
		} else {
			result, err = dao.NewCommentDaoInstance().DeleteComment(commentId)
		}
	}
	return result, err
}

// CommentToCommentInfo 将Comment组装成CommentInfo
func (c CommentServiceImpl) CommentToCommentInfo(comment *entity.Comment, info *entity.CommentInfo, userId int64) {

	userInfo := NewUserServiceImplInstance().GetUserInfo(userId, comment.UserId)
	info.CreateDate = comment.CreateDate.String()
	info.Content = comment.CommentText
	info.UserInfo = userInfo
	info.Id = comment.Id
	log.Println("CommentInfo组装完成")
}

func (c CommentServiceImpl) GetCommentById(commentId int64) (entity.Comment, error) {
	comment, err := dao.NewCommentDaoInstance().FindComment(commentId)
	return comment, err
}

func (c CommentServiceImpl) GetCommentsByVideoId(videoId int64, userId int64) ([]entity.CommentInfo, error) {
	commentList, err := dao.NewCommentDaoInstance().CommentListByVideoId(videoId)
	commentInfoList := make([]entity.CommentInfo, len(commentList), len(commentList))
	for index, _ := range commentList {
		c.CommentToCommentInfo(&commentList[index], &commentInfoList[index], userId)
	}
	return commentInfoList, err

}
