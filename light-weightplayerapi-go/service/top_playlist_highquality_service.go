package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type TopPlaylistHighqualityService struct {
	Limit    string `json:"limit" form:"limit"`
	Cat      string `json:"cat" form:"cat"`
	LastTime string `json:"lasttime" form:"lasttime"`
}

func (service *TopPlaylistHighqualityService) TopPlaylistHighquality(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Cat == "" {
		service.Cat = "全部"
	}
	if service.Limit == "" {
		service.Limit = "50"
	}
	if service.LastTime == "" {
		service.LastTime = "0"
	}
	data["limit"] = service.Limit
	data["lasttime"] = service.LastTime
	data["total"] = "true"
	data["cat"] = service.Cat

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/highquality/list`, data, options)

	return reBody
}
