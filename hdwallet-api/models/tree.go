package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionArticle holds the name of the articles collection
	CollectionTree = "trees"
)

type Tree struct {
    ID          bson.ObjectId   `json:"id" bson:"_id"`
    Coin        string          `json:"coin" bson:"coin"`
    Key         string          `json:"key" bson:"key"`
    ChildCount  int             `json:"childcount" bson:"childcount"`
}
