package main

import (
    //"net/http"
    //"encoding/json"
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    //"database/sql"
     _ "github.com/go-sql-driver/mysql"
     "strconv"
     "crypto/rand"
    //"app/db"
    //"app/controller
)

func init() {
	//db.Connect()
}

type User struct {
    gorm.Model
    Email       string  `json:"email"`
    Password    string  `json:"password"`
    FirstName   string  `json:"first_name"`
    LastName    string  `json:"last_name"`
    Account     Account
}

type Account struct {
    gorm.Model
    Name        string  `json:"name"`
    Address1    string  `json:"address1"`
    Address2    string  `json:"address2"`
    City        string  `json:"city"`
    State       string  `json:"state"`
    Postal      string  `json:"postal"`
    CountryID   uint    `json:"country_id"`
    UserID      uint    `json:"user_id"`
    Merchants   []Merchant
}

type Country struct {
    gorm.Model
    Name        string
    Code        string
    Accounts    []Account
}

type Merchant struct {
    gorm.Model
    Name                string              `json:"name"`
    URL                 string              `json:"url"`
    Token               string              `json:"token"`
    Email               string              `json:"email"`
    PaymentWindow       int                 `json:"payment_window"`
    PaymentThreshold    float64             `json:"payment_threshold"`
    AccountID           uint                `json:"account_id"`
    Transactions        []Transaction
    MerchantCoins       []MerchantCoin
    ApiKeys             []ApiKey
}

type ApiKey struct {
    gorm.Model
    Key         string
    MerchantID  uint
}

type Transaction struct {
    gorm.Model
    PublicKey           string
    Address             string
    CoinTicker          string
    OrderID             string
    AmountDueCoin       float64
    AmountDueUSD        float64
    AmountReceivedCoin  float64
    DateReceivedCoin    time.Time
    MerchantID          uint
    CoinID              uint
}

type MerchantCoin struct {
    gorm.Model
    MerchantID         uint
    CoinID             uint
    DepositAddress     string
}

type Coin struct{
    gorm.Model
    Name                string              `json:"name"`
    Ticker              string              `json:"ticker"`
    MerchantCoins       []MerchantCoin
}

// type MerchantCurrency struct {
//     gorm.Model
//     MerchantID          uint
//     CoinID              uint
//     CurrencyExchangeID  uint
//     DepositAddress      string
//     Transactions        []Transaction
// }
//
// type Exchange struct {
//     gorm.Model
//     Name                string
//     URL                 string
//     ApiToken            string
//     CurrencyExchanges   []CurrencyExchange
//     MerchantCurrencies  []MerchantCurrency
// }
//
// type CurrencyExchange struct {
//     gorm.Model
//     CoinID      uint
//     ExchangeID  uint
//     RateApiUrl  string
// }


// middleware
func ApiKeyAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        validateToken(c)
        c.Next()
    }
}

func validateToken(c *gin.Context) {
    token := c.Request.Header.Get("X-Auth-Token")

    if token == "" {
        c.AbortWithStatus(401)
        return
    } else if checkToken(token) {
        c.Next()
    } else {
        c.AbortWithStatus(401)
        return
    }
}

func checkToken(token string) bool{
    if token == "3sdf33fsf3fx55343v" {
        return true
    } else {
        return false
    }
}

func DummyMiddleware(c *gin.Context) {
  fmt.Println("Im a dummy!")

  // Pass on to the next-in-chain
  c.Next()
}

//Utilities
func randToken() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}



func main () {

    //Database connection, setup and seeding
    db, err := gorm.Open("mysql", "root:secret@tcp(172.20.0.3:3306)/payment?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        fmt.Printf("Can't connect to mysql, go error %v\n", err)
		panic(err.Error())

	}
    defer db.Close()

    db.AutoMigrate(&User{}, &Account{}, &Country{}, &Merchant{}, &Transaction{}, &MerchantCoin{}, &Coin{})

    //seed db
    var cnt int
    db.Table("coins").Count(&cnt)
    fmt.Println("count: ", cnt)
    if cnt == 0 {
        coin := Coin{Name: "Bitcoin", Ticker: "BTC"}
        db.Create(&coin)
        // exchange := Exchange{Name: "Coinbase", URL: "coinbase.com", ApiToken: "23qwefqwefw"}
        // db.Create(&exchange)
        // fmt.Println(coin.ID)
        // currencyExchange := CurrencyExchange{CoinID: coin.ID, ExchangeID: exchange.ID, RateApiUrl: "http://api.coinbase.com/api"}
        // db.Create(&currencyExchange)
        db.Create(&Country{Name:"United States", Code:"US"})
    }

    r := gin.Default()
    // Global middleware
    // Logger middleware will write the logs to gin.DefaultWriter even you set with GIN_MODE=release.
    // By default gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())

    //r.Use(ApiKeyAuthMiddleware())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
			"message": "pong",
		})
    })

    authorized := r.Group("/")

    authorized.Use(ApiKeyAuthMiddleware())
    {

        authorized.GET("/pingauth", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })

    }

    r.POST("/user", func(c *gin.Context) {
        var json User
        c.BindJSON(&json)
        fmt.Println("json data ", json)
        user := User{Email: json.Email, Password: json.Password, FirstName: json.FirstName, LastName: json.LastName}
        db.Create(&user)


        c.JSON(200, gin.H{
            "message": "created",
        })
    })

    r.POST("/account/:userid", func(c *gin.Context) {
        userID := c.Params.ByName("userid")
        i, err := strconv.ParseUint(userID, 10, 64)

        if err == nil {
            var json Account
            c.BindJSON(&json)
            fmt.Println("json data ", json)
            account := Account{Name: json.Name, Address1: json.Address1, Address2: json.Address2, City: json.City, State: json.State, Postal: json.Postal, CountryID: json.CountryID, UserID: uint(i)}
            db.Create(&account)

            c.JSON(200, gin.H{
                "message": "merchant created",
            })
        } else {
            c.JSON(500, gin.H{
                "error": "user id nan",
            })
        }
    })

    r.POST("/merchant/:accountid", func(c *gin.Context) {
        accountid := c.Params.ByName("accountid")
        i, err := strconv.ParseUint(accountid, 10, 64)

        if err == nil {
            var json Merchant
            c.BindJSON(&json)
            fmt.Println("json data ", json)

            merchant := Merchant{Name: json.Name, URL: json.URL, Token: randToken(), Email: json.Email, PaymentWindow: json.PaymentWindow, PaymentThreshold: json.PaymentThreshold, AccountID: uint(i)}
            db.Create(&merchant)

            c.JSON(200, gin.H{
                "message": "merchant created",
            })
        } else {
            c.JSON(500, gin.H{
                "error": "user id nan",
            })
        }
    })


    r.Run(":8080")
}
