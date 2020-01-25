package main

import (
    "net/http"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "app/btcwallet"
    "app/db"
    "app/controllers"
)

func init() {
	db.Connect()
}

func main() {
    r := gin.Default()
    // Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

    // Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin":"top@life",
	}))

    authorized.GET("/ping", func(c *gin.Context) {
        seed := btcwallet.GenerateRandomMnemonicSeed()
        c.JSON(200, gin.H{
			"message": seed,
		})
    })

    treeCtrl := new(controllers.TreeController)
    r.GET("/createtree/:coin", treeCtrl.New)

    merchantCtrl := new(controllers.MerchantController)
    r.GET("/createmerchant/:coin/:token", merchantCtrl.New)
    r.GET("/createmerchantkey/:coin/:token", merchantCtrl.NewKey)


    // r.GET("/createtree/:coin", func(c *gin.Context) {
    //     coin := c.Params.ByName("coin")
    //     m, xprv, xpub  := btcwallet.GenerateRandomMnemonicAndKeys("loveit")
    //
    //     databaseCreateTree(coin, xpub)
    //
    //     c.JSON(200, gin.H{
	// 		"mnemonic": m,
    //         "xprv": xprv,
    //         "xpub": xpub,
	// 	})
    // })

    //
    // r.GET("/createmerchant/:coin/:token", func(c *gin.Context) {
    //     coin := c.Params.ByName("coin")
    //     token := c.Params.ByName("token")
    //     databaseCreateMerchant(coin, token)
    //     c.String(http.StatusOK, "ok")
    // })

    // r.GET("/createmerchantkey/:coin/:token", func(c *gin.Context) {
    //     coin := c.Params.ByName("coin")
    //     token := c.Params.ByName("token")
    //     a := databaseGetMerchantAddress(coin, token)
    //
    //     json, _ := json.Marshal(a)
    //     fmt.Println("json", string(json))
    //     c.String(http.StatusOK, string(json))
    // })

    r.GET("/generateaddress/:coin/:key", func(c *gin.Context) {
        // coin := c.Params.ByName("coin")
        key := c.Params.ByName("key")
        addresses := btcwallet.GenerateChildren(0, 2, key)
        fmt.Println(addresses)
        json, _ := json.Marshal(addresses)
        c.String(http.StatusOK, string(json))
    })

    r.GET("/genchildkeyfromtoken/:token", func(c *gin.Context) {
        publicKey := c.Params.ByName("token")

        addresses := btcwallet.GenerateChildren(0, 2, publicKey)
        fmt.Println(addresses)
        json, _ := json.Marshal(addresses)
        c.String(http.StatusOK, string(json))
    })

    r.Run(":8080")
}

// -----------------------------------------------
// Database
// -----------------------------------------------
//
// func databaseCreateTree(coin string, key string) {
//     c := db.Session.DB("hdwallet").C("tree")
//     err := c.Insert(&models.Tree{bson.NewObjectId(), coin, key, 0})
//     if err != nil {
//         log.Fatal(err)
//         fmt.Println("did not work")
//     }
//
//     fmt.Println("tree created")
// }
//
// func databaseGetTree(coin string) *models.Tree {
//     c := db.Session.DB("hdwallet").C("tree")
//
//     result := models.Tree{}
//     err := c.Find(bson.M{"coin": coin}).One(&result)
//     if err != nil {
//         log.Fatal(err)
//         fmt.Println("nothing found")
//     }
//
//     return &result
// }
//
// func databaseUpdateTree(tree *models.Tree) {
//     c := db.Session.DB("hdwallet").C("tree")
//     err := c.UpdateId(tree.ID, tree)
//     if err != nil {
//         log.Fatal(err)
//         fmt.Println("error updateing tree ", err)
//     }
// }
//
// func databaseGetMerchant(treeID bson.ObjectId, token string) *models.Merchant{
//     result := models.Merchant{}
//
//     c := db.Session.DB("hdwallet").C("merchant")
//     err := c.Find(bson.M{"treeid": treeID, "token": token}).One(&result)
//     if err != nil {
//         log.Fatal(err)
//         fmt.Println("nothing found")
//     }
//     return &result
// }
//
// func databaseUpdateMerchant(merchant *models.Merchant) {
//     c := db.Session.DB("hdwallet").C("merchant")
//     err := c.UpdateId(merchant.ID, merchant)
//     if err != nil {
//         log.Fatal(err)
//         fmt.Println("error updateing tree ", err)
//     }
// }
//
// func databaseCreateMerchant(coin string, token string) {
//     tree := databaseGetTree(coin)
//     fmt.Println(tree.Coin)
//
//     a := btcwallet.GenerateChild(tree.ChildCount, tree.Key)
//
//     fmt.Println("addkey ", a)
//
//     //TODO fix this
//     newCount := tree.ChildCount + 1
//     tree.ChildCount = newCount
//
//     databaseUpdateTree(tree)
//
//     c := db.Session.DB("hdwallet").C("merchant")
//     err := c.Insert(&models.Merchant{bson.NewObjectId(), token, a.Key, 0, tree.ID})
//     if err != nil {
//         log.Fatal(err)
//         fmt.Println("error creating merchant ", err)
//     }
// }
//
// func databaseGetMerchantAddress(coin string, token string) models.Address{
//     tree := databaseGetTree(coin)
//     fmt.Println(tree.Coin)
//
//     merchant := databaseGetMerchant(tree.ID, token)
//
//     a := btcwallet.GenerateChild(merchant.ChildCount, merchant.Key)
//     fmt.Println("addkey ", a)
//
//     //TODO fix this
//     newCount := merchant.ChildCount + 1
//     merchant.ChildCount = newCount
//
//     databaseUpdateMerchant(merchant)
//
//     fmt.Println("get merchant address worked")
//
//     return a
// }
