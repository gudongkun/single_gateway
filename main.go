package main

import (
	"github.com/gudongkun/single_common/custom_gorm"
	"github.com/gudongkun/single_common/jaeger"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"single_gateway/global"
)

func main() {
	//初始化jaeger
	_, cl, _ := jaeger.NewJaegerTracer(global.Conf.App, global.Conf.Jaeger)
	defer cl.Close() //必须close否则不上传
	//初始化 gorm
	custom_gorm.InitEngine(global.Conf.Dsn)
	//初始化 go micro 客户端
	enlight_ucenter_client.InitClient(global.Conf.Service.UCenter)
	//初始化http
	global.RunGin()
	//r := gin.Default()
	//r.POST("/ping", func(c *gin.Context) {
	//	// 初始化span
	//	rootCtx, rootSpan, _ := wrapperTrace.StartSpanFromContext(
	//		c,
	//		opentracing.GlobalTracer(),
	//		c.Request.Method + ":"+c.FullPath(),
	//	)
	//	defer rootSpan.Finish()
	//
	//
	//
	//
	//	c.Set("rootCtx", rootCtx)
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//	body, _ := ioutil.ReadAll(c.Request.Body)
	//	jaeger_log.Info(
	//		rootCtx,
	//		"common_info",
	//		log.String("method", c.Request.Method ),
	//		log.String("path", c.Request.RequestURI),
	//		log.String("body",  string(body)),
	//		//log.String("return", c.Writer.),
	//	)
	//})
	//r.Run()

}
