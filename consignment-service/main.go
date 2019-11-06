package main

import (
	"fmt"
	pb "github.com/codershore/microsrv/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
	"log"
	"os"
)

const defaultHost  = "localhost:27017"

func main()  {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	//Create a new service. Optionally include some options
	srv := micro.NewService(
		micro.Name("consignment"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session})

	//Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

