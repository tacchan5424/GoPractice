package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"../service"
)

type Controller struct{}

func (pc Controller) SearchUserInfo(c *gin.Context) {
	// パラメータの取得
	appName := c.Query("appName")
	//userId := c.Param("userId")
	searchDiv := c.Query("searchDiv")

	// アプリ名からIDとパスワードを取得
	p, err := service.GetUserInfoFromAppName(appName, searchDiv)

  // 区分に応じてサービスコール
	// if searchDiv == "1" {
	// 	p, err := service.GetUserInfoFromAppName(c.Param(appName), c.Param(searchDiv))
	// } else if searchDiv == "2" {
	// 	p, err := service.GetUserInfoFromAppName(c.Param(userId), c.Param(searchDiv))
	// } else {
	// 	p, err := service.GetAllUserInfo(c.Param(searchDiv))
	// }

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
  } else {
		fmt.Println(p)
		c.JSON(200, p)
  }
}