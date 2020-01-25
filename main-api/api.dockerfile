FROM golang:1.9.2

RUN mkdir -p /go/src/app

WORKDIR /go/src/app
COPY . .

RUN go get github.com/jinzhu/gorm && go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql 

RUN go build main.go

CMD ["/go/src/app/main"]
