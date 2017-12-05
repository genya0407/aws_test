package main

import (
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
    "github.com/genya0407/aws/utils"
)

func basicAuth(c *gin.Context) {
    username, password, ok := c.Request.BasicAuth()

    if ok == true && username == "amazon" && password == "candidate" {
        c.String(200, "SUCCESS")
    } else {
        c.Header("WWW-Authenticate", `Basic realm=""`)
        c.String(401, "FORBIDDEN")
    }
}

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(gzip.Gzip(gzip.DefaultCompression))
    r.GET("/secret/", basicAuth)

    return r
}

func main() {
    port := utils.LookupOrDefaultEnv("80", "PORT")

    r := setupRouter()
    r.Run(":" + port)
}