# protobuf Serializer

## test environment

```bash
docker-compose -f ../docker-compose.yml up -d etcd nats
docker-compose -f ../docker-compose.yml down
```

## proto gen

```bash

```

## start servers and tests

```bash
# frontend server
go run .

# backend
go run . -port=3251 -type=room -frontend=false

# client for test
go run . -cli=true

```
