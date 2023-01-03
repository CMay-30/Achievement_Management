package routers

import (
	"manage_backend/controller"

	"github.com/gin-gonic/gin"
)

//路由设置
func RegisterRouter(router *gin.Engine) {
	routerUser(router)
}

//用户路由
func routerUser(engine *gin.Engine) {
	var group = engine.Group("/api/user")
	{
		con := &controller.UserController{}
		group.POST("/add", con.Add)
		group.GET("/get", con.Get)
	}
}
