package service

type CommentService interface {

	// CountFromVideoId 查询视频评论数量
	CountFromVideoId(videoId int64) (int64, error)
}
