# cluster chat


## test environment

```bash
docker-compose -f ../docker-compose.yml up -d etcd nats redis
docker-compose -f ../docker-compose.yml down
```

## star cluster

```bash
# start cluster
go run .
go run . -type worker
go run . -type log
```

## test

> 可以启动多个客户端测试

```bash
# test with cli
cd ../pitaya-cli
go run .

connect 127.0.0.1:3250
request game.account.login {"username":"123123","password":"123123"}
request game.room.join
request game.room.message {"msg":"Hello 123"}
```