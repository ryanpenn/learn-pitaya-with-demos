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
