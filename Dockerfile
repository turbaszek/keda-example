FROM golang:1.15

WORKDIR /go/src/app
COPY helper/ .

RUN go get -v ./...

RUN go install -v .

ENTRYPOINT ["/go/bin/keda-talk"]
