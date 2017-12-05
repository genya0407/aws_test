package handler

import (
    "strconv"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/genya0407/aws/stocker"
)

type StockerHandler struct {
    Stocker stocker.Stocker
}

func (sh *StockerHandler) Handle(c *gin.Context) {
    switch c.Query("function") {
    case "addstock":
        sh.addStock(c)
    /*
    case "checkstock":
        sh.checkStock(c)
    case "sell":
        sh.sell(c)
    case "checksales":
        sh.checkSales(c)
    case "deleteall":
        sh.deleteAll(c)
    */
    default:
        log.Println("invalid function!")
        c.String(400, "ERROR")
    }
    return
}

func (sh *StockerHandler) addStock(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
        c.String(400, "ERROR")
        return
    }
    amount, err := strconv.Atoi(c.DefaultQuery("amount", "1"))
    if err != nil {
        c.String(400, "ERROR")
        return
    }

    log.Println("before addstock")
    sh.Stocker.AddStock(name, amount)
    log.Println("after addstock")
    c.String(200, "")
}