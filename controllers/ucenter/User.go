package ucenter

import (
	"github.com/gin-gonic/gin"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/client/user"
	"github.com/prometheus/common/log"
)

func GetUser(c *gin.Context) {
	res, err  := user.GetName(c, 1)
	if err != nil {
		log.Info(err)
	}
	user.GetAge(c, 1)
	user.GetSex(c, 1)
	c.JSON(200, res)
}
