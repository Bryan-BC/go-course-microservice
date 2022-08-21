proto:
	del /S *.pb.go
	protoc pkg/pb/course.proto --go_out=plugins=grpc:.

start:
	go run cmd/main.go
	