package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Router 路由转发
func Router(r *gin.Engine) {
	fmt.Println("Router...")
	r.GET("/ping", helloWorld)
	r.GET("/get_details", getAccountDetails)
	r.POST("/modify_account", modifyAccount)
}
