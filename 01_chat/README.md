# chat-demo

> chat room demo base on [pitaya](https://github.com/topfreegames/pitaya)

## Required
- golang
- websocket
- docker

## Run
```
docker-compose -f ../docker-compose.yml up -d etcd nats
go run main.go

docker-compose -f ../docker-compose.yml down
```

open browser => http://localhost:3251/
