# cluster-demo

## test environment

```bash
docker-compose -f ../docker-compose.yml up -d etcd nats
docker-compose -f ../docker-compose.yml down
```

## frontend server start

```bash
go run main.go
```

## backend server start

```bash
go run main.go -port=3251 -type=room -frontend=false
```

## test with client

```bash
cd ../pitaya-cli
go run .

connect 127.0.0.1:3250

# call backend server
request room.room.entry
request room.room.join

# call frontend server
request connector.setsessiondata {"Data":{"ipversion":"ipv4","notice":"some message"}}
request connector.getsessiondata

notify connector.setsessiondata {"Data":{"ipversion":"ipv4","notice":"a new message"}}
request connector.getsessiondata
```