package main

import (
	"github.com/gudongkun/single_common/custom_gorm"
	"github.com/gudongkun/single_common/jaeger"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"single_gateway/global"
	"single_gateway/globalgin"
)

func main() {
	//初始化jaeger
	_, cl, _ := jaeger.NewJaegerTracer(global.Conf.App, global.Conf.Jaeger)
	defer cl.Close() //必须close否则不上传
	//初始化 gorm
	custom_gorm.InitEngine(global.Conf.Dsn)
	//初始化 casbin
	global.InitCasbin(global.Conf.CasbinDsn)
	//初始化 go micro 客户端
	enlight_ucenter_client.InitClient(global.Conf.Service.UCenter)
	//初始化http
	globalgin.RunGin()


}
