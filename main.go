package main

import (
	"douyin-backend/util"
	"github.com/gin-gonic/gin"
)

func main() {
	util.Init()
	r := gin.Default()
	initRouter(r)
	r.Run("localhost:8080")

}
