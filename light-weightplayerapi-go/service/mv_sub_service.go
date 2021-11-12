package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type MvSubService struct {
	T    string `json:"t" form:"t"`
	MvId string `json:"mvid" form:"mvid"`
}

func (service *MvSubService) MvSub(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.T == "1" {
		service.T = "sub"
	} else {
		service.T = "unsub"
	}

	data["mvId"] = service.MvId
	data["mvIds"] = "[" + service.MvId + "]"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/mv/`+service.T, data, options)

	return reBody
}
