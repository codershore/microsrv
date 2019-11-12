build:
	protoc --proto_path="/Users/funtech/go/pkg/mod/github.com/micro/go-micro@v1.15.1/api" \
	 --go_out=. --micro_out=. proto/api.proto

# 	GOOS=linux GOARCH=amd64 go build
# 	docker build -t shippy-service-user .


# run:
# 	docker run -p 50053:50051 \
# 		-e MICRO_SERVER_ADDRESS=:50051 \
# 		-e MICRO_REGISTRY=mdns \
# 		shippy-service-user