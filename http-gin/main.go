package main

import (
	"github.com/gin-gonic/gin"

	"http-gin/apis"
)

func main() {
	// 默认启动方式，包含 Logger、Recovery 中间件
	r := gin.Default()
	/*
		TODO
		1、路由分组
		2、中间件
		3、写日志文件
		4、模型绑定和验证
		5、
	*/
	apis.Router(r)
	r.Run()
}
