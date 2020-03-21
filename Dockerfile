FROM danisbagus/base-go:gomod

RUN mkdir -p /go

COPY app /go

EXPOSE 80

WORKDIR /go

RUN rm go.mod go.sum

RUN go mod init go-project

RUN go get github.com/joho/godotenv
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm

CMD ["go","run", "main.go"]
