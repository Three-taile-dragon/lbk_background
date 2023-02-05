package main

import (
	_ "dragonsss.cn/lbk_api/api"
	"dragonsss.cn/lbk_api/config"
	"dragonsss.cn/lbk_api/router"
	srv "dragonsss.cn/lbk_common"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(logs.GinLogger(), logs.GinRecovery(true)) //接收gin框架默认日志
	//注册路由
	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
