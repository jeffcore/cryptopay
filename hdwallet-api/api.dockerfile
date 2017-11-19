FROM golang:1.9.2

RUN mkdir -p /go/src/app

WORKDIR /go/src/app
COPY . .

RUN go get github.com/WeMeetAgain/go-hdwallet && go get github.com/tyler-smith/go-bip39 && \
    go get github.com/tyler-smith/go-bip32 && go get github.com/gin-gonic/gin && \
    go get gopkg.in/mgo.v2

RUN go build main.go

CMD ["/go/src/app/main"]
