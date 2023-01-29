package dao

import (
	"douyin-backend/entity"
	"douyin-backend/util"
	"log"
	"sync"
)

// FollowDao 把dao层看成整体，把dao的curd封装在一个结构体中。
type LikeDao struct {
}

var (
	likeDao  *LikeDao  //操作该dao层crud的结构体变量。
	likeOnce sync.Once //单例限定，去限定申请一个followDao结构体变量。
)

// NewLikeDaoInstance 生成并返回followDao的单例对象。
func NewLikeDaoInstance() *LikeDao {
	likeOnce.Do(
		func() {
			likeDao = &LikeDao{}
		})
	return likeDao
}

// GetLikeInfo 根据userId,videoId查询点赞信息
func (*LikeDao) GetLikeInfo(userId int64, videoId int64) (entity.Like, error) {
	//创建一条空like结构体，用来存储查询到的信息
	var likeInfo entity.Like
	//根据userid,videoId查询是否有该条信息，如果有，存储在likeInfo,返回查询结果
	err := util.DB.Model(&entity.Like{}).
		Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		First(&likeInfo).Error
	if err != nil {
		//查询数据为0，打印"can't find data"，返回空结构体，这时候就应该要考虑是否插入这条数据了
		if "record not found" == err.Error() {
			log.Println("查不到对应点赞关系")
			return likeInfo, nil
		} else {
			//如果查询数据库失败，返回获取likeInfo信息失败
			log.Println(err.Error())
			return likeInfo, err
		}
	}
	return likeInfo, nil
}

// GetLikeUserIdList 根据videoId获取点赞userId
func (*LikeDao) GetLikeUserIdList(videoId int64) ([]int64, error) {
	var likeUserIdList []int64 //存所有该视频点赞用户id；
	//查询likes表对应视频id点赞用户，返回查询结果
	err := util.DB.Model(&entity.Like{}).
		Where(map[string]interface{}{"video_id": videoId, "cancel": 0}).
		Pluck("user_id", &likeUserIdList).Error
	//查询过程出现错误，返回默认值0，并输出错误信息
	if err != nil {
		log.Println("获取视频点赞列表出错")
		return nil, err
	} else {
		//没查询到或者查询到结果，返回数量以及无报错
		return likeUserIdList, nil
	}
}
