package apis

import (
	"github.com/gin-gonic/gin"
)

func modifyAccount(c *gin.Context) {
	//获取参数
	acccountID := c.PostForm("account_id")
	balance := c.DefaultPostForm("balance", "10")

	//@ TODO修改玩家数据，db操作

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "modify account info successfully",
	})
}
