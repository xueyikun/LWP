package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type CloudsearchService struct {
	S      string `json:"keywords" form:"keywords"`
	Type   string `json:"type" form:"type"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *CloudsearchService) Cloudsearch(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Type == "" {
		service.Type = "1"
	}
	if service.Limit == "" {
		service.Limit = "30"
	}
	if service.Offset == "" {
		service.Offset = "0"
	}
	data["limit"] = service.Limit
	data["offset"] = service.Offset
	data["type"] = service.Type
	data["s"] = service.S

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/cloudsearch/pc`, data, options)

	return reBody
}
