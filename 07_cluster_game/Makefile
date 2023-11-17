ifeq ($(OS),Windows_NT)
RM = del /s /q /f
else
RM = rm -rf
endif

.PHONY: protos
# 生成协议
protos:
	@echo "step1: clean protos folder"
	@cd protos
	$(RM) *.pb.go
	@cd ..
	@echo "step2: gen *.pb.go files"
	@protoc --go_out=. -I=./assets/protos  ./assets/protos/*.proto
	@echo "step2: done!"

# 启动容器
docker-up:
	docker-compose -f ../docker-compose.yml up -d nats etcd mongo

# 停止容器
docker-down:
	docker-compose -f ../docker-compose.yml down

.PHONY: login
# login-server
login:
	go run . start login
