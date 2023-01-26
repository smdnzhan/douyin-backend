package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserLoginResponse struct {
	Response entity.Response
	UserId   int64  `json:"user_id,omitempty"`
	Token    string `json:"token"`
}

// Register POST douyin/user/register/ 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	usi := service.UserServiceImpl{}

	u := usi.GetUserPOByName(username)
	if username == u.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//暂时不加密
		newUser := entity.UserPO{
			Name:     username,
			Password: password,
		}
		if usi.InsertUserPO(&newUser) != true {
			println("Insert Data Fail")
		}
		//注册成功 返回id和token
		u := usi.GetUserPOByName(username)
		//token暂不加密
		token := username
		log.Println("注册返回的id: ", u.Id)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   u.Id,
			Token:    token,
		})
	}
}

// Login POST douyin/user/login/ 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := password
	println(encoderPassword)

	usi := service.UserServiceImpl{}

	u := usi.GetUserPOByName(username)

	if encoderPassword == u.Password {
		token := username
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   u.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Username or Password Error"},
		})
	}
}
