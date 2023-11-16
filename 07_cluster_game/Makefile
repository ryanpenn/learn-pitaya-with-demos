ifeq ($(OS),Windows_NT)
RM = rmdir /s /q
else
RM = rm -rf
endif

.PHONY: protos
# 生成协议
protos:
	@echo "step1: clean pd folder"
	$(RM) protos/*
	@echo "step2: gen *.pb.go files"
	@protoc --go_out=. -I=./assets/protos  ./assets/protos/*.proto
	@echo "step2: done!"
