package handler

import "github.com/gin-gonic/gin"

func BasicAuth(c *gin.Context) {
    username, password, ok := c.Request.BasicAuth()

    if ok == true && username == "amazon" && password == "candidate" {
        c.String(200, "SUCCESS")
    } else {
        c.Header("WWW-Authenticate", `Basic realm=""`)
        c.String(401, "FORBIDDEN")
    }
}