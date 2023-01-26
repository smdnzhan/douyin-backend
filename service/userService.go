package service

import "douyin-backend/entity"

type userService interface {
	//DAO层查询结果 暂定返回UserPO
	GetUserList() []entity.UserPO
	GetUserPOById(id int64) entity.UserPO
	GetUserPOByName(name string) entity.UserPO
	InsertUserPO(userPO *entity.UserPO) bool
}
