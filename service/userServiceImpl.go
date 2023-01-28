package service

import (
	"douyin-backend/dao"
	"douyin-backend/entity"
	"douyin-backend/util"
	"log"
	"sync"
)

type UserServiceImpl struct {
}

var (
	userServiceImpl *UserServiceImpl //controller层通过该实例变量调用service的所有业务方法。
	userServiceOnce sync.Once        //限定该service对象为单例，节约内存。
)

// NewUserServiceImplInstance 生成并返回followServiceImpl的单例对象。
func NewUserServiceImplInstance() *UserServiceImpl {
	userServiceOnce.Do(
		func() {
			userServiceImpl = &UserServiceImpl{}
		})
	return userServiceImpl
}

// GetUserPOById 根据id获得UserPO对象
func (usi *UserServiceImpl) GetUserPOById(id int64) entity.UserPO {
	userPO, err := dao.NewUserDaoInstance().GetUserPOById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return userPO
	}
	log.Println("Query User Success")
	return userPO
}

func (usi *UserServiceImpl) GetUserPOByName(name string) entity.UserPO {
	userPO, err := dao.NewUserDaoInstance().GetUserPOByName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return userPO
	}
	log.Println("Query User Success")
	return userPO
}

func (usi *UserServiceImpl) GetUserList(ids []int64) []entity.UserPO {
	userPOs, err := dao.NewUserDaoInstance().GetUserPOByIds(ids)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return userPOs
	}
	log.Println("Query User Success")
	return userPOs
}

func (usi *UserServiceImpl) InsertUserPO(userPO *entity.UserPO) bool {
	result, err := dao.NewUserDaoInstance().InsertUserPO(userPO)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return result
	}
	log.Println("Query User Success")
	return result
}

func (usi *UserServiceImpl) GenerateToken(username string) (string, error) {
	userPO, err := dao.NewUserDaoInstance().GetUserPOByName(username)
	if err != nil {
		log.Println("生成token时出现问题")
		return "", err
	} else {
		token := util.NewToken(userPO)
		return token, err
	}

}

// UNGetUserInfo 用于未登录用户获取其他用户关注+粉丝列表
func (usi *UserServiceImpl) UNGetUserInfo(id int64) entity.UserInfo {
	follower_cnt, err := dao.NewFollowDaoInstance().GetFollowerCnt(id)
	if err != nil {
		log.Println("获取用户粉丝数出错：", err.Error())
	}
	follow_cnt, err := dao.NewFollowDaoInstance().GetFollowingCnt(id)
	if err != nil {
		log.Println("获取用户关注数出错：", err.Error())
	}
	//返回结果里不能带密码，所以新建一个User对象
	userPO, err := dao.NewUserDaoInstance().GetUserPOById(id)
	userInfo := entity.UserInfo{
		Id:            userPO.Id,
		Name:          userPO.Name,
		FollowCount:   follow_cnt,
		FollowerCount: follower_cnt,
		IsFollow:      false,
	}
	return userInfo
}

// GetUserInfo 用于已登录用户获取其他用户关注+粉丝列表
func (usi *UserServiceImpl) GetUserInfo(userId int64, targetId int64) entity.UserInfo {

	follower_cnt, err := NewFollowServiceImpInstance().GetFollowerCnt(targetId)
	if err != nil {
		log.Println("获取用户粉丝数出错：", err.Error())
	}
	follow_cnt, err := NewFollowServiceImpInstance().GetFollowingCnt(targetId)
	if err != nil {
		log.Println("获取用户关注数出错：", err.Error())
	}

	is_follow, err := NewFollowServiceImpInstance().IsFollow(userId, targetId)
	if err != nil {
		log.Println("获取用户关系出错：", err.Error())
	}
	var flag bool
	if is_follow == nil {
		flag = false
	}
	if is_follow.Cancel == 0 {
		flag = true
	}
	if is_follow.Cancel == 1 {
		flag = false
	}
	//返回结果里不能带密码，所以新建一个User对象
	userPO, err := dao.NewUserDaoInstance().GetUserPOById(targetId)
	userInfo := entity.UserInfo{
		Id:            userPO.Id,
		Name:          userPO.Name,
		FollowCount:   follow_cnt,
		FollowerCount: follower_cnt,
		IsFollow:      flag,
	}
	return userInfo
}
