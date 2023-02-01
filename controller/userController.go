package controller

import (
	"douyin-backend/entity"
	"douyin-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Register POST douyin/user/register/ 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	usi := service.NewUserServiceImplInstance()

	u := usi.GetUserPOByName(username)
	if username == u.Name {
		c.JSON(http.StatusOK, entity.UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User already exist",
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
		c.JSON(http.StatusOK, entity.UserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "Success",
			UserId:     u.Id,
			Token:      token,
		})
	}
}

// Login POST douyin/user/login/ 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := password
	println(encoderPassword)

	usi := service.NewUserServiceImplInstance()

	u := usi.GetUserPOByName(username)

	if encoderPassword == u.Password {
		token, error := usi.GenerateToken(username)
		if error == nil {
			c.JSON(http.StatusOK, entity.UserLoginResponse{
				StatusCode: 0,
				StatusMsg:  "Success",
				UserId:     u.Id,
				Token:      token,
			})
		} else {
			c.JSON(http.StatusOK, entity.UserLoginResponse{
				StatusCode: 1,
				StatusMsg:  "fail",
				UserId:     0,
				Token:      "",
			})
		}
	} else {
		c.JSON(http.StatusOK, entity.UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "Username or Password Error",
		})
	}
}

func UserInfo(c *gin.Context) {
	//目标用户Id
	targetId := c.Query("user_id")
	targetIdInt, _ := strconv.ParseInt(targetId, 10, 64)
	log.Println("目标用户id：" + targetId)
	//从上下文中获取 当前用户id
	userId := c.GetString("user_id")
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	log.Println("当前用户id：" + userId)
	usi := service.NewUserServiceImplInstance()
	var result entity.UserInfo
	//当前用户是未登录用户，查询其他用户信息
	if len(userId) == 0 {
		log.Println("当前用户是未登录用户")
		result = usi.UNGetUserInfo(targetIdInt)
	} else {
		log.Println("当前用户是登录用户")
		result = usi.GetUserInfo(userIdInt, targetIdInt)
	}
	c.JSON(http.StatusOK, entity.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
		UserInfo:   result,
	})
}
