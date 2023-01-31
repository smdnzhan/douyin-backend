package dao

import (
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
	log.Println("关注数:", cnt)
	// 查询成功，返回人数。
	return cnt, nil
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
		log.Println(err.Error())
		return &follow, err
	}
	log.Println("follow查询结果:", follow)
	//正常情况，返回取到的值和空err.
	return &follow, nil
}
