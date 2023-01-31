package main

import (
	"douyin-backend/controller"
	"douyin-backend/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	// 基础功能
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.FormAuth(), controller.Publish)
	//中间件判断用户是否登录
	apiRouter.GET("/user/", middleware.QueryAuth(), controller.UserInfo)
	apiRouter.GET("/publish/list/", middleware.QueryAuth(), controller.PublishList)
<<<<<<< HEAD
	apiRouter.GET("/feed/", middleware.LoginOrNot(), controller.Feed)

	//扩展功能
	apiRouter.POST("/favorite/action/", middleware.QueryAuth(), controller.Favorite)
	apiRouter.GET("/favorite/list/", middleware.QueryAuth(), controller.FavoriteList)
=======
>>>>>>> c2e33fd9cbe428c8b1809cbece6f06d4b70dde3d
}
