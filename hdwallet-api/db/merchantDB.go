package db

import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2/bson"
    "app/models"
    "app/btcwallet"
)

//ArticleController ...
type MerchantDB struct{}

var merchantModel = new(models.Merchant)

func (merchantDB MerchantDB) Create(coin string, token string) {

    treeDB := new(TreeDB)

    tree := treeDB.Get(coin)
    fmt.Println(tree.Coin)

    a := btcwallet.GenerateChild(tree.ChildCount, tree.Key)

    fmt.Println("addkey ", a)

    //TODO fix this mady should fix this
    //TODO mady needs to learn how to program
    newCount := tree.ChildCount + 1
    tree.ChildCount = newCount

    treeDB.Update(tree)

    c := Session.DB("hdwallet").C("merchant")
    err := c.Insert(&models.Merchant{bson.NewObjectId(), token, a.Key, 0, tree.ID})
    if err != nil {
        log.Fatal(err)
        fmt.Println("error creating merchant ", err)
    }
}

func (merchantDB MerchantDB) Get(treeID bson.ObjectId, token string) *models.Merchant{
    result := models.Merchant{}

    c := Session.DB("hdwallet").C("merchant")
    err := c.Find(bson.M{"treeid": treeID, "token": token}).One(&result)
    if err != nil {
        log.Fatal(err)
        fmt.Println("nothing found")
    }
    return &result
}

func (merchantDB MerchantDB) Update(merchant *models.Merchant) {
    c := Session.DB("hdwallet").C("merchant")
    err := c.UpdateId(merchant.ID, merchant)
    if err != nil {
        log.Fatal(err)
        fmt.Println("error updateing tree ", err)
    }
}

func (merchantDB MerchantDB) CreateAddress(coin string, token string) models.Address{
    treeDB := new(TreeDB)
    tree := treeDB.Get(coin)
    fmt.Println(tree.Coin)

    merchant := MerchantDB{}.Get(tree.ID, token)

    a := btcwallet.GenerateChild(merchant.ChildCount, merchant.Key)
    fmt.Println("addkey ", a)

    //TODO fix this
    newCount := merchant.ChildCount + 1
    merchant.ChildCount = newCount

    MerchantDB{}.Update(merchant)

    fmt.Println("get merchant address worked")

    return a
}
