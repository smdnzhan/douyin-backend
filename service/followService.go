package service

import "douyin-backend/entity"

type FollowService interface {

	// GetFollowerCnt 根据用户id来查询用户的粉丝数量
	GetFollowerCnt(userId int64) (int64, error)
	// GetFollowingCnt 根据用户id来查询用户关注了多少其它用户
	GetFollowingCnt(userId int64) (int64, error)
	IsFollow(userId int64, targetId int64) (*entity.Follow, error)
	UpdateFollow(userId int64, targetId int64) error
}
