FROM golang:alpine as build

RUN apk add --no-cache git 

WORKDIR /go/app

COPY . .

ENV GO111MODULE=on

RUN go mod download

RUN go get -u github.com/thaitanloi365/govvv && \
    GO111MODULE=off go get github.com/erikdubbelboer/realize

ENV GO111MODULE=on

FROM alpine

WORKDIR /app

COPY --from=build /go/app .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]