package main

import (
	"github.com/gin-gonic/gin"
	"manage_backend/config"
	"manage_backend/routers"
	"manage_backend/utils"
)

func main() {
	conf, err := config.ParseConfig("./config/app.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	//fmt.Printf("conf:%#v\n", conf)

	utils.InitRedisUtil(conf.RedisConfig.Addr, conf.RedisConfig.Port, conf.RedisConfig.Password)

	//创建一个默认的路由引擎
	engine := gin.Default()
	routers.RegisterRouter(engine)
	engine.Run(":9090")
}
