build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/user/user.proto

	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-user .


run:
	docker run -p 50053:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		shippy-service-user