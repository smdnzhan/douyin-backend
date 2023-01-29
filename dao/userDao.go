package dao

import (
	"douyin-backend/entity"
	"douyin-backend/util"
	"log"
	"sync"
)

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

// NewUserDaoInstance 返回UserDAO单例
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) InsertUserPO(userDO *entity.UserPO) (bool, error) {
	err := util.DB.Create(&userDO).Error
	if err != nil {
		log.Println("新增用户失败:" + err.Error())
		return false, err
	} else {
		return true, err
	}
}

func (*UserDao) GetUserPOByName(name string) (entity.UserPO, error) {
	var result entity.UserPO
	err := util.DB.Where("name= ?", name).Find(&result).Error
	if err != nil {
		log.Println("使用名称查询用户失败:" + err.Error())
	}
	return result, err

}
func (*UserDao) GetUserPOById(id int64) (entity.UserPO, error) {
	var result entity.UserPO
	err := util.DB.Where("id= ?", id).Find(&result).Error
	if err != nil {
		log.Println("使用id查询用户失败:" + err.Error())
	}
	return result, err
}

// 使用id切片查询 返回用户DO切片
func (*UserDao) GetUserPOByIds(ids []int64) ([]entity.UserPO, error) {
	var result []entity.UserPO
	err := util.DB.Where("id= ?", ids).Find(&result).Error
	if err != nil {
		log.Println("使用id切片查询用户失败:" + err.Error())
	}
	return result, err
}
