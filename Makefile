proto:
	C:\protoc-21.5-win64\bin\protoc pkg/pb/course.proto --go_out=plugins=grpc:.

start:
	go run cmd/main.go
	