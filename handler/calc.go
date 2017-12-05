package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/genya0407/aws/parser"
    "strings"
    "fmt"
)

const validChars = "+-*/()0123456789"
func CalcHandler(c *gin.Context) {
    query := c.Request.URL.RawQuery
    for i := 0; i < len(query); i++ {
        if !strings.Contains(validChars, string(query[i])) {
            c.String(200, "ERROR")
            return
        }
    }

    result := parser.Parse(strings.TrimSpace(query))
    c.String(200, fmt.Sprintf("%d", result))
    return
}