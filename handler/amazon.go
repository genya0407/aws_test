package handler

import "github.com/gin-gonic/gin"

func Amazon(c *gin.Context) {
    c.String(200, "AMAZON")
}