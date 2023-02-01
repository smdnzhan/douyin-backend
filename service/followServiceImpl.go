package service

import (
	"douyin-backend/dao"
	"douyin-backend/entity"
	"log"

	//"strconv"
	"sync"
)

// FollowServiceImp 该结构体继承FollowService接口。
type FollowServiceImpl struct {
}

var (
	followServiceImp  *FollowServiceImpl //controller层通过该实例变量调用service的所有业务方法。
	followServiceOnce sync.Once          //限定该service对象为单例，节约内存。
)

// NewFollowServiceImpInstance 生成并返回followServiceImpl的单例对象。
func NewFollowServiceImpInstance() *FollowServiceImpl {
	followServiceOnce.Do(
		func() {
			followServiceImp = &FollowServiceImpl{}
		})
	return followServiceImp
}

// GetFollowerCnt 给定当前用户id，查询其粉丝数量。
func (*FollowServiceImpl) GetFollowerCnt(userId int64) (int64, error) {
	//redis..
	// SQL中查询。
	cnt, err := dao.NewFollowDaoInstance().GetFollowerCnt(userId)
	if nil != err {
		return 0, err
	}
	// 将数据存入Redis.
	// 更新followers 和 followingPart
	//go addFollowersToRedis(int(userId), ids)

	return cnt, err
}

// GetFollowingCnt 给定当前用户id，查询其关注者数量。
func (*FollowServiceImpl) GetFollowingCnt(userId int64) (int64, error) {
	// redis...
	// 用SQL查询。
	cnt, err := dao.NewFollowDaoInstance().GetFollowingCnt(userId)

	if nil != err {
		return 0, err
	}
	// 更新Redis中的followers和followPart
	//go addFollowingToRedis(int(userId), ids)

	return cnt, err
}

func (*FollowServiceImpl) IsFollow(userId int64, targetId int64) (*entity.Follow, error) {
	// redis...
	// 用SQL查询。
	follow, err := dao.NewFollowDaoInstance().IsFollow(userId, targetId)
	return follow, err
}

func (*FollowServiceImpl) UpdateFollow(userId int64, targetId int64) error {
	err := dao.NewFollowDaoInstance().UpdateFollow(userId, targetId)
	if err != nil {
		log.Printf("更新关注关系出错")
	}
	return err
}

func (*FollowServiceImpl) CreateFollow(userId int64, targetId int64) error {
	follow, err := dao.NewFollowDaoInstance().IsFollow(userId, targetId)
	log.Println("新增关注:", follow)
	return err
}

func (*FollowServiceImpl) FollowList(userId int64, tagetId int64) ([]entity.UserInfo, error) {
	//首先获取目标用户的关注用户id列表
	following, err := dao.NewFollowDaoInstance().GetFollowingList(tagetId)
	//组装成UserInfo
	var followingUserInfo = make([]entity.UserInfo, len(following), len(following))
	if userId != 0 {
		for index, _ := range following {
			followingUserInfo[index] = NewUserServiceImplInstance().GetUserInfo(userId, following[index])
		}
	} else {
		for index, _ := range following {
			followingUserInfo[index] = NewUserServiceImplInstance().UNGetUserInfo(following[index])
		}
	}
	return followingUserInfo, err
}
