build:
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 api/proto/v1/*.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 api/proto/v1/*.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 api/proto/v1/*.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/server/ ./cmd/server
