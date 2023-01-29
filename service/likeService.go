package service

type LikeService interface {

	//IsFavorite 根据当前视频id判断是否点赞了该视频。
	IsFavorite(videoId int64, userId int64) (bool, error)
	//FavouriteCount 根据当前视频id获取当前视频点赞数量。
	FavouriteCount(videoId int64) (int64, error)
}

