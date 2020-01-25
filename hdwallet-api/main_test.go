package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "app/models"
    "app/btcwallet"
    "github.com/stretchr/testify/assert"
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

func TestPingRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "pong", w.Body.String())
}
