package service

import "douyin-backend/entity"

type UserService interface {
	// GetUserList DAO层查询结果 暂定返回UserPO
	GetUserList(ids []int64) []entity.UserPO
	GetUserPOById(id int64) entity.UserPO
	GetUserPOByName(name string) entity.UserPO
	InsertUserPO(userPO *entity.UserPO) bool
	GenerateToken(username string) (string, error)

	// UNGetUserInfo 向上层返回包装对象UserInfo
	UNGetUserInfo(id int64) entity.UserInfo
	GetUserInfo(userId int64, targetId int64) entity.UserInfo
}
