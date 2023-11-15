# msgp Serializer

> 自定义序列化器

## test environment

```bash
docker-compose -f ../docker-compose.yml up -d etcd nats
docker-compose -f ../docker-compose.yml down
```

## run and test

```bash
# server
go run .

# client
go run . -cli=true

```