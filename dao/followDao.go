package dao

import (
	"douyin-backend/config"
	"douyin-backend/entity"
	"douyin-backend/util"
	"log"
	"sync"
)

// FollowDao 把dao层看成整体，把dao的curd封装在一个结构体中。
type FollowDao struct {
}

var (
	followDao  *FollowDao //操作该dao层crud的结构体变量。
	followOnce sync.Once  //单例限定，去限定申请一个followDao结构体变量。
)

// NewFollowDaoInstance 生成并返回followDao的单例对象。
func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

// GetFollowerCnt 给定当前用户id，查询follow表中该用户的粉丝数。
func (*FollowDao) GetFollowerCnt(userId int64) (int64, error) {
	// 用于存储当前用户粉丝数的变量
	var cnt int64
	// 当查询出现错误的情况，日志打印err msg，并返回err.
	if err := util.DB.
		Model(entity.Follow{}).
		Where("user_id = ?", userId).
		Where("cancel = ?", 0).
		Count(&cnt).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	// 正常情况，返回取到的粉丝数。
	log.Println("粉丝数:", cnt)
	return cnt, nil
}

// GetFollowingCnt 给定当前用户id，查询follow表中该用户关注了多少人。
func (*FollowDao) GetFollowingCnt(userId int64) (int64, error) {
	// 用于存储当前用户关注了多少人。
	var cnt int64
	// 查询出错，日志打印err msg，并return err
	if err := util.DB.
		Model(entity.Follow{}).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Count(&cnt).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	// 查询成功，返回人数。
	return cnt, nil
}

// 返回用户当前
func (*FollowDao) GetFollowingList(userId int64) ([]int64, error) {
	// 用于存储当前用户关注的人。
	var cnt []entity.Follow
	// 查询出错，日志打印err msg，并return err
	var result []int64
	if err := util.DB.
		Model(entity.Follow{}).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Find(&cnt).Error; nil != err {
		log.Println(err.Error())
		return result, err
	}
	result = make([]int64, len(cnt), len(cnt))
	for index, _ := range cnt {
		result[index] = cnt[index].UserId
	}
	log.Println("关注用户切片:", result)
	// 查询成功，返回被关注人id切片。
	return result, nil
}

func (*FollowDao) IsFollow(userId int64, targetId int64) (*entity.Follow, error) {
	// follow变量用于后续存储数据库查出来的用户关系。
	follow := entity.Follow{}
	//当查询出现错误时，日志打印err msg，并return err.
	if err := util.DB.
		Where("user_id = ?", targetId).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Take(&follow).Error; nil != err {
		// 当没查到数据时，gorm也会报错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		return &follow, err
	}
	//正常情况，返回取到的值和空err.
	return &follow, nil
}

// 查找是否有历史关注关系
func (*FollowDao) ExistFollow(userId int64, targetId int64) (*entity.Follow, error) {
	// follow变量用于后续存储数据库查出来的用户关系。
	follow := entity.Follow{}
	//当查询出现错误时，日志打印err msg，并return err.
	if err := util.DB.
		Where("user_id = ?", targetId).
		Where("follower_id = ?", userId).
		Take(&follow).Error; nil != err {
		// 当没查到数据时，gorm也会报错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		return &follow, err
	}
	//正常情况，返回取到的值和空err.
	return &follow, nil
}

func (*FollowDao) InsertFollow(userId int64, toUserId int64) error {
	follow := entity.Follow{
		UserId:     toUserId,
		FollowerId: userId,
		Cancel:     config.FOLLOW_STATUS,
	}
	err := util.DB.Model(follow).Create(&follow)
	if err.Error != nil {
		log.Println("新增关注出错", err.Error)
	}
	return err.Error
}

func (*FollowDao) UpdateFollow(userId int64, toUserId int64) error {
	follow, _ := NewFollowDaoInstance().ExistFollow(userId, toUserId)
	var err error
	if follow == nil {
		log.Println("查询不到历史关注关系")
		err = NewFollowDaoInstance().InsertFollow(userId, toUserId)

	} else {
		log.Println("查询到的关注关系:", follow)
		status := follow.Cancel
		if status == 1 {
			log.Println("恢复关注")
			follow.Cancel = 0
		} else {
			log.Println("取消关注")
			follow.Cancel = 1
		}
		err = util.DB.Model(follow).
			Where("id = ?", follow.Id).
			Update("cancel", follow.Cancel).Error

	}
	return err

}
