package handler

import (
    "strconv"
    "log"
    "fmt"
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
    case "checkstock":
        sh.checkStock(c)
    case "sell":
        sh.sell(c)
    /*
    case "checksales":
        sh.checkSales(c)
    */
    case "deleteall":
        sh.deleteAll(c)
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
    if err != nil || amount <= 0 {
        c.String(400, "ERROR")
        return
    }

    sh.Stocker.AddStock(name, amount)
    c.String(200, "")
}

func (sh *StockerHandler) checkStock(c *gin.Context) {
    name := c.Query("name")
    items, _ := sh.Stocker.CheckStock(name)

    var responseBody string
    for _, item := range items {
        responseBody = responseBody + item.Name + ": " + fmt.Sprintf("%d", item.Amount) + "\n"
    }
    c.String(200, responseBody)
}

func (sh *StockerHandler) sell(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
        c.String(400, "ERROR")
        return
    }
    amount, err := strconv.Atoi(c.DefaultQuery("amount", "1"))
    if err != nil || amount <= 0 {
        c.String(400, "ERROR")
        return
    }
    price, err := strconv.Atoi(c.DefaultQuery("price", "0"))
    if err != nil || amount < 0 {
        c.String(400, "ERROR")
        return
    }

    sh.Stocker.Sell(name, amount, price)
    c.String(200, "")
}

func (sh *StockerHandler) deleteAll(c *gin.Context) {
    sh.Stocker.DeleteAll()
    c.String(200, "")
}