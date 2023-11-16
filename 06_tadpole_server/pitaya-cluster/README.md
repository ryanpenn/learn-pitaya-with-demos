# Tadpole

## test environment

```bash
docker-compose -f ../../docker-compose.yml up -d etcd nats
docker-compose -f ../../docker-compose.yml down
```

## Run demo
```shell
# frontend
go run .

# backend
go run . -front=false

```

Open browser: http://127.0.0.1:9000/static/