package ucenter

import (
	"github.com/gin-gonic/gin"
	"github.com/gudongkun/single_common/jaeger_log"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/client/user"
	log2 "github.com/opentracing/opentracing-go/log"
	"github.com/prometheus/common/log"
)

func GetUser(ctx *gin.Context) {
	c:= ctx.Request.Context()  //特别注意jaeger不支持gin的context
	res, err  := user.GetName(c, 1)

	jaeger_log.Info(c,"test",log2.String("hhh", "hhhhhhh"),)
	if err != nil {
		log.Info(err)
	}
	user.GetAge(c, 1)
	user.GetSex(c, 1)
	ctx.JSON(200, res)
}
