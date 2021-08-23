package globalgin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"single_gateway/global"
	"single_gateway/middlewares"
	"single_gateway/router"
)

func GetRouter() *gin.Engine {
	var Router = gin.Default()

	//Router.NoMethod(middleware.HandleNotFound)
	//Router.NoRoute(middleware.HandleNotFound)
	//// 跨域
	//Router.Use(middleware.Cors())
	//Router.Use(middleware.ErrHandler())
	//Router.Use(middleware.TpsCheckHandler())
	//注册中间件
	Router.Use(middlewares.AutoJaeger())
	//注册路由
	rootGroup := Router.Group("")
	router.InitUCenterRouter(rootGroup)

	log.Info("run router success")
	return Router
}

func RunGin()  {
	router := GetRouter()
	router.Run()
	if err := router.Run(global.Conf.AppAddr); err != nil {
		log.Error("http Service Start Error", err)
	}
}
