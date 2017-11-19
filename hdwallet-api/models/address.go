package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionAddress = "address"
)

type Address struct {
	Key         string // 33 bytes
	Addr     string // 4 bytes
}
