FROM golang:1.17-alpine

WORKDIR /commentify

ADD ./app ./app
ADD ./base ./base
ADD ./business ./business
ADD ./vendor ./vendor
ADD go.mod ./go.mod
ADD go.sum ./go.sum

CMD ["go", "run", "app/server/main.go"]