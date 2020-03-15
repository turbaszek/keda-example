FROM golang:1.13

WORKDIR /go/src/app
COPY helper .

RUN go get github.com/go-redis/redis github.com/urfave/cli

RUN go install

ENTRYPOINT ["app"]