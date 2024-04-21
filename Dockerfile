FROM golang:1.20.4 as builder

RUN go version
LABEL authors="enyad"

COPY ./ /go/src/github.com/enyaaad/ztrback/
WORKDIR /go/src/github.com/enyaaad/ztrback/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

RUN apk update
RUN apk add git
RUN apk upgrade
RUN apk add postgresql-client
RUN apk --no-cache add curl

COPY --from=0  /go/src/github.com/enyaaad/ztrback/.bin/app .
COPY --from=0  /go/src/github.com/enyaaad/ztrback/config/ ./config/

CMD ["./app"]

