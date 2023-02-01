package entity

// Follow 用户关系结构，对应用户关系表。
type Follow struct {
	Id         int64
	UserId     int64
	FollowerId int64
	Cancel     int8
}

func (follow Follow) TableName() string {
	return "follows"
}

// UserInfoResponse 用户信息响应
type UserInfoListResponse struct {
	StatusCode   int32      `json:"status_code"`
	StatusMsg    string     `json:"status_msg,omitempty"`
	UserInfoList []UserInfo `json:"user_list"`
}
