package main

import (
	"douyin-backend/util"
	"github.com/gin-gonic/gin"
)

func main() {
	util.Init()
	util.InitOSS()
	r := gin.Default()
	initRouter(r)
	r.Run(":8080")

}
