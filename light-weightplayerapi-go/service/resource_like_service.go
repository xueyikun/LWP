package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type ResourceLikeService struct {
	ID       string `json:"id" form:"id"`
	ThreadId string `json:"threadId" form:"threadId"`
	T        string `json:"t" form:"t"`
	Type     string `json:"type" form:"type"`
}

func (service *ResourceLikeService) ResourceLike(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	TYPE := make(map[string]string, 6)
	TYPE["1"] = "R_MV_5_"
	TYPE["4"] = "A_DJ_1_"
	TYPE["5"] = "R_VI_62_"
	TYPE["6"] = "A_EV_2_"

	if _, ok := TYPE[service.Type]; ok {
		service.Type = TYPE[service.Type]
	} else {
		service.Type = TYPE["1"]
	}

	data := make(map[string]string)
	data["threadId"] = service.Type + service.ID

	if service.Type == "A_EV_2_" {
		data["threadId"] = service.ThreadId
	}

	if service.T == "1" {
		service.T = "like"
	} else {
		service.T = "unlike"
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/resource/`+service.T, data, options)

	return reBody
}
