package service

import (
	"douyin-backend/dao"
	"log"
	"sync"
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

<<<<<<< HEAD
// 查询评论数量
func (c CommentServiceImpl) CommentCountFromVideoId(videoId int64) (int64, error) {
=======
func (c CommentServiceImpl) CountFromVideoId(videoId int64) (int64, error) {
>>>>>>> c2e33fd9cbe428c8b1809cbece6f06d4b70dde3d
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
