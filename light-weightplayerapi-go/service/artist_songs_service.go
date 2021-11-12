package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type ArtistSongsService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
	Order  string `json:"order" form:"order"`
}

func (service *ArtistSongsService) ArtistSongs(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	if service.Limit == "" {
		data["limit"] = "100"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	if service.Order == "" {
		data["order"] = "hot"
	} else {
		data["order"] = service.Order
	}
	data["work_type"] = "1"
	data["private_cloud"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v1/artist/songs`, data, options)

	return reBody
}
