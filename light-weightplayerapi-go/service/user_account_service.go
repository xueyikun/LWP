package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserAccountService struct {
}

func (service *UserAccountService) UserAccount(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/api/nuser/account/get`, data, options)

	return reBody
}
