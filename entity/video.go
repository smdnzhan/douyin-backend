package entity

type Video struct {
	Id       int64 `json:"id,omitempty"`
	AuthorId int64
	PlayUrl  string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl string `json:"cover_url,omitempty"`
}

type VideoInfo struct {
	Video         Video
	Author        User  `json:"author"`
	FavoriteCount int64 `json:"favorite_count"`
	CommentCount  int64 `json:"comment_count"`
	IsFavorite    bool  `json:"is_favorite"`
}
