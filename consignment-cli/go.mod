module github.com/codershore/micorsrv/consignment-cli

go 1.13

// replace github.com/codershore/micorsrv/consignment-service/proto/consignment => ../consignment-service/proto/consignment

require (
	github.com/codershore/microsrv/consignment-service v0.0.0-20191030002141-1e8e1ccdf43c
	github.com/micro/go-micro v1.14.0
)
