package main
import (
    "testing"
    "app/models"
    "app/btcwallet"
)

func TestPing(t *testing.T) {
  // Test some functionality in here
}


func TestTwoPlusTwo(t *testing.T) {
    twoPlusTwo := 2 + 2
    if twoPlusTwo != 5 {
        t.Error("Expected 2 + 2 to equal 5, but got", twoPlusTwo)
    }
}


//btcwallet test

func TestGenerateRandomMnemonicSeed(t *testing.T) {
    seed := btcwallet.GenerateRandomMnemonicSeed()

    if seed == nil {
        t.Error("nothign returned")
    }    
}
