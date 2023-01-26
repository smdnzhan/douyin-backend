package service

import (
	"douyin-backend/dao"
	"douyin-backend/entity"
	"log"
)

type UserServiceImpl struct {
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

func (usi *UserServiceImpl) getUserList(ids []int64) []entity.UserPO {
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
