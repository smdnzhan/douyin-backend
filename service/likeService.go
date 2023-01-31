package service

<<<<<<< HEAD
import "douyin-backend/entity"

type LikeService interface {

	//IsFavorite 根据当前视频id判断是否点赞了该视频。
	IsFavorite(userId int64, videoId int64) (bool, error)
	//FavouriteCount 根据当前视频id获取当前视频点赞数量。
	FavouriteCount(videoId int64) (int64, error)

	UpdateLike(videoId int64, userId int64, status string) error
	GetLikedVideoList(userId int64) ([]entity.VideoInfo, error)
}
=======
type LikeService interface {

	//IsFavorite 根据当前视频id判断是否点赞了该视频。
	IsFavorite(videoId int64, userId int64) (bool, error)
	//FavouriteCount 根据当前视频id获取当前视频点赞数量。
	FavouriteCount(videoId int64) (int64, error)
}

>>>>>>> c2e33fd9cbe428c8b1809cbece6f06d4b70dde3d
