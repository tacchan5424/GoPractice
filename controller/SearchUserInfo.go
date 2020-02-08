package controller

import (
	"github.com/gin-gonic/gin"
	"../entity"
	"../service"
)

func SearchUserInfo(c *gin.Context, appName string, userId string, searchDiv string) []entity.SearchResult {
	if searchDiv == "1" {
		return service.GetUserInfoFromAppName(c.Param(appName), c.Param(searchDiv))
	} else if searchDiv == "2" {
		return service.GetUserInfoFromAppName(c.Param(userId), c.Param(searchDiv))
	} else {
		return service.GetAllUserInfo(c.Param(searchDiv))
	}
}