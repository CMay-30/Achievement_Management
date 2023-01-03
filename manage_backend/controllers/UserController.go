package controller

import (
	"manage_backend/config"
	"manage_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

//新增用户
func (controller *UserController) Add(context *gin.Context) {
	name, exist := context.GetPostForm("name")
	if !exist || name == "" {
		context.JSON(http.StatusOK, gin.H{
			"msg": "请输入用户名:name",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

//查询用户
func (controller *UserController) Get(context *gin.Context) {
	id := context.Query("id")
	utils.RedisClient.SetStr("hello", "111")
	utils.RedisClient.SetStrWithExpire("hello22", "22", 100)
	hello := utils.RedisClient.GetStr("hello")
	context.JSON(http.StatusOK, gin.H{
		"id":    id,
		"conf":  config.GetConfig(),
		"hello": hello,
	})
}
