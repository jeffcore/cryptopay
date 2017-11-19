package main

import (
    "github.com/tyler-smith/go-bip39"
    "github.com/tyler-smith/go-bip32"
    "github.com/WeMeetAgain/go-hdwallet"
    "app/models"
)

func GenerateRandomMnemonicSeed() string {
    entropy, _ := bip39.NewEntropy(256)
    mnemonic, _ := bip39.NewMnemonic(entropy)

    return mnemonic
}

func GenerateRandomMnemonicAndKeys(passPhrase string) (string, string, string) {
    entropy, _ := bip39.NewEntropy(256)
    mnemonic, _ := bip39.NewMnemonic(entropy)
    seed := bip39.NewSeed(mnemonic, passPhrase)
    masterKey, _ := bip32.NewMasterKey(seed)
    publicKey := masterKey.PublicKey()

    return mnemonic, masterKey.String(), publicKey.String()
}

func GenerateChildren(startIndex int, totalCount int, publicKey string) []models.Address {
    //figure out how to create the data structure
    addresses := []Address{}

    for i := startIndex; i < startIndex+totalCount; i++ {
        childstring, _ := hdwallet.StringChild(publicKey,uint32(i))
        addr, _ := hdwallet.StringAddress(childstring)
        a := Address{Key: childstring , Addr: addr }
        addresses = append(addresses,a)
    }

    return addresses
}

func GenerateChild(startIndex int, publicKey string) model.Address {
    childstring, _ := hdwallet.StringChild(publicKey,uint32(startIndex))
    addr, _ := hdwallet.StringAddress(childstring)
    a := Address{Key: childstring , Addr: addr }
    return a
}



// func database() {
//     session, err := mgo.Dial("mongodb://hdwallet-mongo:27017")
//     if err != nil {
//         panic(err)
//     }
//     defer session.Close()
//
//     c := session.DB("hdwallet").C("merchant")
//     err = c.Insert(&Merchant{bson.NewObjectId(),
//         "bitcore2",
//         "xpub661MyMwAqRbcFSMJFEniyNUZrhrUtpBRMjuWzfgAtoygvFMUZKXi9jPxMJzfoTLZ8VhrtXkDk9JZVtVMF3TjfchsXf3jRufdbXz1aMgVxVP",
//         0,
//         "ObjectId('5a09fd22054f39dea1f23f0f')",
//     })
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     result := Merchant{}
//     err = c.Find(bson.M{"token": "bitcore"}).One(&result)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     fmt.Println("Merchant:", result.Token)
// }

// func databaseGenKeyFromToken(token string) address{
//     session, err := mgo.Dial("mongodb://hdwallet-mongo:27017")
//     if err != nil {
//         panic(err)
//     }
//     defer session.Close()
//
//     result := merchant{}
//     err = c.Find(bson.M{"token": "bitcore"}).One(&result)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//
//     c := session.DB("hdwallet").C("merchant")
//     err = c.Insert(&merchant{"bitcore2", "xpub661MyMwAqRbcFSMJFEniyNUZrhrUtpBRMjuWzfgAtoygvFMUZKXi9jPxMJzfoTLZ8VhrtXkDk9JZVtVMF3TjfchsXf3jRufdbXz1aMgVxVP", 0})
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     result := merchant{}
//     err = c.Find(bson.M{"token": "bitcore"}).One(&result)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     fmt.Println("Merchant:", result.token)
// }
