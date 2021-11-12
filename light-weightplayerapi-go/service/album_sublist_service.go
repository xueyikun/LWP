package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type AlbumSublistService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *AlbumSublistService) AlbumSublist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "25"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["total"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/album/sublist`, data, options)

	return reBody
}
