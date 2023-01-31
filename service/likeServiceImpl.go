package service

import (
	"douyin-backend/dao"
	"douyin-backend/entity"
	"log"
	"sync"
)

type LikeServiceImpl struct {
}

var (
	likeServiceImpl *LikeServiceImpl //controller层通过该实例变量调用service的所有业务方法。
	likeServiceOnce sync.Once        //限定该service对象为单例，节约内存。
)

// NewLikeServiceImplInstance 生成并返回LikeServiceImpl的单例对象。
func NewLikeServiceImplInstance() *LikeServiceImpl {
	likeServiceOnce.Do(
		func() {
			likeServiceImpl = &LikeServiceImpl{}
		})
	return likeServiceImpl
}

// IsFavorite 根据当前视频id判断是否点赞了该视频。
func (*LikeServiceImpl) IsFavorite(userId int64, videoId int64) (bool, error) {
	var like entity.Like
	like, err := dao.NewLikeDaoInstance().GetLikeInfo(userId, videoId)
	if err != nil {
		//查不到的逻辑
		log.Printf("查询点赞关系出错", err.Error())
		return false, err
	} else if like.Id == 0 || like.Cancel == 1 {
		//空结构体意味着没有历史点赞关系 Cancel为1为先点赞后取消点赞的情况
		return false, nil
	} else {
		return true, nil
	}
}

// FavouriteCount 根据当前视频id获取当前视频点赞数量。
func (*LikeServiceImpl) FavouriteCount(videoId int64) (int64, error) {
	var userlist []int64
	userlist, err := dao.NewLikeDaoInstance().GetLikeUserIdList(videoId)
	return int64(len(userlist)), err
}

func (*LikeServiceImpl) UpdateLike(videoId int64, userId int64, status string) error {
	var err error
	//取消点赞必定已经点赞过 则更新关系
	if status == "2" {
		log.Println("取消点赞")
		err = dao.NewLikeDaoInstance().UpdateLike(videoId, userId)
		//未点赞过 或取消点赞而恢复
	} else {
		log.Println("点赞")
		err = dao.NewLikeDaoInstance().UpdateLike(videoId, userId)
	}

	return err
}

func (*LikeServiceImpl) GetLikedVideoList(userId int64) ([]entity.VideoInfo, error) {
	videoIds, err := dao.NewLikeDaoInstance().LikedVideoList(userId)
	if err != nil {
		log.Println("查询用户点赞视频出错")
	}
	resultList, err := NewVideoServiceImplInstance().GetVideoListByIds(videoIds)
	result := make([]entity.VideoInfo, len(videoIds), len(videoIds))
	for index, temp := range resultList {
		NewVideoServiceImplInstance().VideoPOToVideoInfo(&temp, &result[index], userId)
	}
	return result, err

}
