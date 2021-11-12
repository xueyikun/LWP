package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 获取歌单分类
func PlaylistCatlist(c *gin.Context) {
	var service service.PlaylistCatlistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistCatlist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热门歌单分类
func PlaylistHot(c *gin.Context) {
	var service service.PlaylistHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistHot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌单详情
func PlaylistDetail(c *gin.Context) {
	var service service.PlaylistDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// 收藏歌单
func PlaylistSubscribe(c *gin.Context) {
	var service service.PlaylistSubscribeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistSubscribe(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌单收藏者
func PlaylistSubscribers(c *gin.Context) {
	var service service.PlaylistSubscribersService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistSubscribers(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
