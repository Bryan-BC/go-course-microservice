# go-course-microservice

This is a microservice that handles the creation and retrieval of a course. This microservice is built in conjunction with the following:

- [API Gateway](https://github.com/Bryan-BC/go-api-gateway)
- [Auth Service](https://github.com/Bryan-BC/go-auth-microservice)
- [Timetable Service](https://github.com/Bryan-BC/go-timetable-microservice)

## Prerequisites

Since this service utilises gRPC and protocol buffers, it requires the [protocol buffer compiler](https://grpc.io/docs/protoc-installation/) installed. Moreover, it also uses `Makefile` to compile the protobuf files. This means it also requires [make](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows) to be installed

## Setup

To run the service, simply cd into the repository and run the following command:

`make start`

Once this is run, the gateway will be running on port 5000.
