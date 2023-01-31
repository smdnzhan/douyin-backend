package dao

import (
<<<<<<< HEAD
	"douyin-backend/config"
=======
>>>>>>> c2e33fd9cbe428c8b1809cbece6f06d4b70dde3d
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
<<<<<<< HEAD
			log.Println("查询数据库失败，返回获取likeInfo信息失败")
=======
			log.Println(err.Error())
>>>>>>> c2e33fd9cbe428c8b1809cbece6f06d4b70dde3d
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
<<<<<<< HEAD

// InsertLike 对没有历史记录视频和用户插入一条关系
func (*LikeDao) InsertLike(videoId int64, userId int64) error {
	relation := entity.Like{
		UserId:  userId,
		VideoId: videoId,
		Cancel:  config.LIKE_STATUS,
	}
	//gorm框架记得用地址
	//reflect: reflect.Value.SetInt using unaddressable value
	//	这种报错是应该传地址的时候传了值导致的
	err := util.DB.Create(&relation).Error
	if err != nil {
		log.Println("添加点赞信息出错")
	}

	return err
}

// UpdateLike 恢复点赞或取消点赞
func (*LikeDao) UpdateLike(videoId int64, userId int64) error {
	var relation entity.Like
	err := util.DB.Model(relation).
		Where("user_id = ? AND video_id = ?", userId, videoId).
		First(&relation).Error
	if err != nil && err.Error() == "record not found" {
		log.Println("新建点赞关系")
		err = NewLikeDaoInstance().InsertLike(videoId, userId)
	} else {
		//恢复点赞或取消点赞
		log.Println("更新点赞关系")
		if relation.Cancel == config.UNLIKE_STATUS {
			err = util.DB.Model(relation).
				Where("user_id = ? AND video_id = ?", userId, videoId).
				Update("cancel", config.LIKE_STATUS).Error
		} else {
			err = util.DB.Model(relation).
				Where("user_id = ? AND video_id = ?", userId, videoId).
				Update("cancel", config.UNLIKE_STATUS).Error
		}

	}
	return err
}

// LikedVideoList 返回用户点赞的视频的集合
func (*LikeDao) LikedVideoList(userId int64) ([]int64, error) {
	var likes []entity.Like
	err := util.DB.Model(entity.Like{}).
		Where(map[string]interface{}{"user_id": userId}).
		Find(&likes).Error
	if err != nil {
		log.Println("查询点赞关系出错")
	}
	result := make([]int64, len(likes), len(likes))
	for index, like := range likes {
		result[index] = like.VideoId
	}
	log.Printf("liked videoList:", result)
	return result, err
}
=======
>>>>>>> c2e33fd9cbe428c8b1809cbece6f06d4b70dde3d
