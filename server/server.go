package server

import (
    "github.com/gin-gonic/gin"

    "github.com/asuforce/gin-gorm-tutorial/controller"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()
    return r
}