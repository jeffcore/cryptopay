package controllers

import (
    "net/http"
    "encoding/json"
    "fmt"
	"github.com/gin-gonic/gin"
	"app/models"
    "app/db"
)


//ArticleController ...
type MerchantController struct{}

var merchantModel = new(models.Merchant)

// New merchant
func (ctrl MerchantController) New(c *gin.Context) {
    coin := c.Params.ByName("coin")
    token := c.Params.ByName("token")

    merchantDB := new(db.MerchantDB)
    merchantDB.Create(coin, token)
    c.String(http.StatusOK, "ok")
}

func (ctrl MerchantController) NewKey(c *gin.Context) {
    coin := c.Params.ByName("coin")
    token := c.Params.ByName("token")

    merchantDB := new(db.MerchantDB)

    a := merchantDB.CreateAddress(coin, token)

    json, _ := json.Marshal(a)
    fmt.Println("json", string(json))
    c.String(http.StatusOK, string(json))
}
