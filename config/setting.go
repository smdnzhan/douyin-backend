package config

const (
	MYSQL_ADDR       = ""
	SECRET_KEY       = "DOUYIN"
	OSS_ENDPOINT     = "oss-cn-hangzhou.aliyuncs.com"
	OSS_ACCESSKEYID  = ""
	OSS_ACCESSSECRET = ""
	OSS_BUCKET       = "zdx-shangcheng"
	//视频存储路径
	VIDEO_PATH = "video"
	//视频封面存储路径
	COVER_PATH = "cover"
	ONEDAY     = 60 * 60 * 24
	//FEED流视频数量
	FEED_COUNT = 5
	// UNLOGIN_USER 未登录用户默认id为0
	UNLOGIN_USER     = 0
	LIKE_STATUS      = 0
	UNLIKE_STATUS    = 1
	COMMENT_STATUS   = 0
	UNCOMMENT_STATUS = 1
	FOLLOW_STATUS    = 0
	UNFOLLOW_STATUS  = 1
)
