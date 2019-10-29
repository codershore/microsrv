module github.com/codershore/microsrv

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.14.0
	google.golang.org/grpc v1.24.0
)

replace (
	"github.com/codershore/microsrv/consignment-service/proto/consignment" => 
	"Users/funtech/go/src/github.com/codershore/microsrv/consignment-service/proto/consignment"
)