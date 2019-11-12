package main

import (
	"context"
	"encoding/json"
	micro "github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"strings"

	proto "github.com/codershore/microsrv/microl/simpleapi/proto/"
)

//Example Service
type Example struct{}

//Foo Service
type Foo struct{}

//Call for Example service
func (e *Example) Call(ctx context.Context, req *api.Request, res *api.Response) error {
	log.Log("Example.Call接口收到请求")

	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "Wrong parameters")
	}

	for k, v := range req.Header {
		log.Log("请求信息， ", k, " : ", v)
	}

	res.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": "w我们已经收到请求，" + strings.Join(name.Values, " "),
	})

	res.Body = string(b)

	return nil
}

//Bar for Foo service
func (f *Foo) Bar(ctx context.Context, req *api.Request, res *api.Response) error {
	log.Log("Foo.Bar接口收到请求")

	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.example", "require post")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "expect application/json")
	}

	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)

	res.Body = "recieve the message: " + string([]byte(req.Body))
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	proto.RegisterExampleHandler(service.Server(), new(Example))

}
