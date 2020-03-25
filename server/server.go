package server

import (
    "github.com/gin-gonic/gin"
    "../controller"
    "github.com/gin-contrib/cors"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    r.Use(cors.New(config))

    u := r.Group("") 
    {
        ctrl := controller.Controller{}
        u.GET("", ctrl.SearchUserInfo)
    }

    return r
}