# 路径
TOP_PATH = $(shell pwd) 
# s 启动服务类型 w-代表web g-代表grpc
s  = w
# 启动服务
run:
	rm -rf $(shell pwd)/protoorder/*.go
	protoc --proto_path=$(shell pwd)/proto/ --micro_out=$(shell pwd)/protoorder --go_out=$(shell pwd)/protoorder $(shell pwd)/proto/*proto
	go run main.go $s
# grpc 客户端测试命令
co:
	go run $(shell pwd)/client/cliord.go
# grpc 客户端测试命令
list:
	go run $(shell pwd)/servicelist/servicelist.go
