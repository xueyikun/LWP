package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 推荐歌单
func Personalized(c *gin.Context) {
	var service service.PersonalizedService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Personalized(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
