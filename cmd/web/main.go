package main

import (
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
    "github.com/genya0407/aws/utils"
    "github.com/genya0407/aws/handler"
    "github.com/genya0407/aws/stocker"
)

func main() {
    port := utils.LookupOrDefaultEnv("80", "PORT")

    db := utils.SetupDB()
    stocker := stocker.Stocker { DB: db }
    stockerHandler := handler.StockerHandler { Stocker: stocker }

    r := gin.Default()
    r.Use(gzip.Gzip(gzip.DefaultCompression))
    r.GET("/", handler.Amazon)
    r.GET("/secret/", handler.BasicAuth)
    r.GET("/calc", handler.CalcHandler)
    r.GET("/stocker", stockerHandler.Handle)

    r.Run(":" + port)
}