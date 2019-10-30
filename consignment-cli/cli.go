package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	// pb "github.com/codershore/micorsrv/consignment-service/proto/consignment"
	pb "github.com/codershore/microsrv/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}
func main() {

	// Create new greeter client
	client := pb.NewShippingServiceClient("consignment", microclient.DefaultClient)
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
