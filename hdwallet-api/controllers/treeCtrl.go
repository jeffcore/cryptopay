package controllers

import (
	"github.com/gin-gonic/gin"
	"app/models"
    "app/db"
    "app/btcwallet"
)


//ArticleController ...
type TreeController struct{}

var treeModel = new(models.Tree)

// New article
func (ctrl TreeController) New(c *gin.Context) {
    coin := c.Params.ByName("coin")
    m, xprv, xpub  := btcwallet.GenerateRandomMnemonicAndKeys("loveit")
    treeDB := new(db.TreeDB)
    treeDB.Create(coin, xpub)

    c.JSON(200, gin.H{
        "mnemonic": m,
        "xprv": xprv,
        "xpub": xpub,
    })
}
