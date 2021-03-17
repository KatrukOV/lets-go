all: clean generateProto

clean:
	find . -name "*.pb.go" -exec rm -rf {} \;

generateProto:
	protoc -I . api/api.proto --go_out=plugins=grpc:.

