package entity

import "time"

// VideoPO  与数据库字段对应
type VideoPO struct {
	Id          int64     `json:"id,omitempty"`
	AuthorId    int64     `gorm:"author_id"`
	PlayUrl     string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl    string    `json:"cover_url,omitempty"`
	PublishTime time.Time `gorm:"publish_time"`
	Title       string    `json:"title"`
}

// VideoInfo 返回给客户端的包装对象
type VideoInfo struct {
	Id            int64    `json:"id,omitempty"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	Title         string   `json:"title"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
}

func (videoPO VideoPO) TableName() string {
	return "videos"
}

type VideoListResponse struct {
	StatusCode    int32       `json:"status_code"`
	StatusMsg     string      `json:"status_msg,omitempty"`
	VideoInfoList []VideoInfo `json:"video_list"`
}
