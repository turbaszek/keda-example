FROM golang:1.13

WORKDIR /go/src/app
COPY helper .

RUN go get \
    github.com/go-redis/redis \
    github.com/urfave/cli \
    github.com/go-sql-driver/mysql

RUN go install

ENTRYPOINT ["app"]