package entity

// UserPO 对应数据库User表结构的结构体
type UserPO struct {
	Id       int64
	Name     string
	Password string
}

func (userPO UserPO) TableName() string {
	return "users"
}

type User struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// UserInfo 在User基础上多了关注/被关注的信息，在其他响应中也会用到
type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserInfoResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg,omitempty"`
	UserInfo   UserInfo `json:"user"`
}
