package main

import (
	srv "dragonsss.cn/lbk_common"
	"dragonsss.cn/lbk_project/config"
	"dragonsss.cn/lbk_project/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(logs.GinLogger(), logs.GinRecovery(true)) //接收gin框架默认日志
	router.InitRouter(r) //路由初始化
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()
	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
