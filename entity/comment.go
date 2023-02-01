package entity

import "time"

// Comment 评论信息结构体-dao层使用
type Comment struct {
	Id          int64     //评论id
	UserId      int64     //评论用户id
	VideoId     int64     //视频id
	CommentText string    //评论内容
	CreateDate  time.Time //评论发布的日期mm-dd
	Cancel      int32     //取消评论为1，发布评论为0
}

// TableName 修改表名映射
func (Comment) TableName() string {
	return "comments"
}

// CommentInfo 返回评论信息
type CommentInfo struct {
	Id         int64 `json:"id,omitempty"`
	UserInfo   `json:"user,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
type CommentInfoResponse struct {
	StatusCode  int32  `json:"status_code"`
	StatusMsg   string `json:"status_msg,omitempty"`
	CommentInfo `json:"comment,omitempty"`
}

type CommentInfoListResponse struct {
	StatusCode      int32         `json:"status_code"`
	StatusMsg       string        `json:"status_msg,omitempty"`
	CommentInfoList []CommentInfo `json:"comment_list,omitempty"`
}
