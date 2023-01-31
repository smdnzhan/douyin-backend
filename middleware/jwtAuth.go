package middleware

import (
	"douyin-backend/config"
	"douyin-backend/entity"
	"douyin-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth 鉴权中间件
// 若用户携带的token正确,解析token,将userId放入上下文context中并放行;否则,返回错误信息
func QueryAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//auth := context.Request.Header.Get("Authorization")
		auth := context.Query("token")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, entity.Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		}

		token, err := util.ParseToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, entity.Response{
				StatusCode: -1,
				StatusMsg:  "Token Error",
			})
		} else {
			println("token 正确,将userId设置进user_id:", token.Id)
		}
		context.Set("user_id", token.Id)
		context.Next()
	}
}

// LoginOrNot 查询用户是否登录 未登录则将0作为user_id放入上下文 会将未登录用户放行
func LoginOrNot() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Query("token")

		if len(auth) == 0 {
			context.Set("user_id", config.UNLOGIN_USER)
		} else {

			token, err := util.ParseToken(auth)
			if err != nil {
				context.Abort()
				context.JSON(http.StatusUnauthorized, entity.Response{
					StatusCode: -1,
					StatusMsg:  "Token Error",
				})
			} else {
				println("token 正确,将userId设置进user_id:", token.Id)
				context.Set("user_id", token.Id)
			}
		}
		context.Next()
	}
}

func FormAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//auth := context.Request.Header.Get("Authorization")
		auth := context.Request.PostFormValue("token")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, entity.Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		}

		claim, err := util.ParseToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, entity.Response{
				StatusCode: -1,
				StatusMsg:  "Token Error",
			})
		} else {
			println("token 正确")
		}
		context.Set("user_id", claim.Id)
		context.Next()
	}
}
