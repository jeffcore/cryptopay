package db

import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2/bson"
    "app/models"
)

type TreeDB struct{}

var treeModel = new(models.Merchant)

func (treeDB TreeDB) Create(coin string, key string) {
    c := Session.DB("hdwallet").C("tree")
    err := c.Insert(&models.Tree{bson.NewObjectId(), coin, key, 0})
    if err != nil {
        log.Fatal(err)
        fmt.Println("did not work")
    }

    fmt.Println("tree created")
}

func (treeDB TreeDB) Get(coin string) *models.Tree {
    c := Session.DB("hdwallet").C("tree")

    result := models.Tree{}
    err := c.Find(bson.M{"coin": coin}).One(&result)
    if err != nil {
        log.Fatal(err)
        fmt.Println("nothing found")
    }

    return &result
}

func (treeDB TreeDB) Update(tree *models.Tree) {
    c := Session.DB("hdwallet").C("tree")
    err := c.UpdateId(tree.ID, tree)
    if err != nil {
        log.Fatal(err)
        fmt.Println("error updateing tree ", err)
    }
}
