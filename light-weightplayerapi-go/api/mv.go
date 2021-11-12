package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 收藏/取消收藏MV
func MvSub(c *gin.Context) {
	var service service.MvSubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MvSub(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 收藏的MV列表
func MvSublist(c *gin.Context) {
	var service service.MvSublistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MvSublist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 全部mv
func MvAll(c *gin.Context) {
	var service service.MvAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MvAll(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取mv详细数据
func MvDetail(c *gin.Context) {
	var service service.MvDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MvDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// mv地址
func MvUrl(c *gin.Context) {
	var service service.MvUrlService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MvUrl(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
