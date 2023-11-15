# chat-demo

> chat room demo base on [pitaya](https://github.com/topfreegames/pitaya)

## Required
- golang
- websocket
- docker

## test environment

```bash
docker-compose -f ../docker-compose.yml up -d etcd nats
docker-compose -f ../docker-compose.yml down
```

## Run
```
go run main.go

# open browser => http://localhost:3251/
```
