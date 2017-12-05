package main

import (
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
    "github.com/genya0407/aws/utils"
    "github.com/genya0407/aws/handler"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(gzip.Gzip(gzip.DefaultCompression))
    r.GET("/", handler.Amazon)
    r.GET("/secret/", handler.BasicAuth)

    return r
}

func main() {
    port := utils.LookupOrDefaultEnv("80", "PORT")

    r := setupRouter()
    r.Run(":" + port)
}