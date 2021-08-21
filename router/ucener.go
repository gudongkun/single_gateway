package router

import (
	"github.com/gin-gonic/gin"
	"single_gateway/controllers/ucenter"
)

func InitUCenterRouter(Router *gin.RouterGroup) {
	//UCenterRouter := Router.Group("UCenter").Use(middleware.JWTAuth())
	UCenterRouter := Router.Group("UCenter")
	{
		UCenterRouter.GET("getUser", ucenter.GetUser) // 修改
	}

}


