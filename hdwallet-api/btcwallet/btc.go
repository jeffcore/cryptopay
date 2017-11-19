package btcwallet

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
    addresses := []models.Address{}

    for i := startIndex; i < startIndex+totalCount; i++ {
        childstring, _ := hdwallet.StringChild(publicKey,uint32(i))
        addr, _ := hdwallet.StringAddress(childstring)
        a := models.Address{Key: childstring , Addr: addr }
        addresses = append(addresses,a)
    }

    return addresses
}

func GenerateChild(startIndex int, publicKey string) models.Address {
    childstring, _ := hdwallet.StringChild(publicKey,uint32(startIndex))
    addr, _ := hdwallet.StringAddress(childstring)
    a := models.Address{Key: childstring , Addr: addr }
    return a
}
