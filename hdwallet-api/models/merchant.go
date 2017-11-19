package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionArticle holds the name of the articles collection
	CollectionMerchant = "merchants"
)

type Merchant struct {
    ID          bson.ObjectId   `json:"id" bson:"_id"`
    Token       string          `json:"token" bson:"token"`
    Key         string          `json:"key" bson:"key"`
    ChildCount  int             `json:"childcount" bson:"childcount"`
    TreeID      bson.ObjectId   `json:"treeid" bson:"treeid"`
}
