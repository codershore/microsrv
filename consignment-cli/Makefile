build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-cli-consignment .

run:
	docker run -e MICRO_REGISTRY=mdns shippy-cli-consignment