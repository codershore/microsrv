package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/codershore/microsrv/consignment-service/proto/consignment"

	"github.com/micro/go-micro"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

//Repository - Dummy repository.
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

//Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

//GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

//CreateConsignment - we created just on method on our service,
//Which is a create method, which takes a context and a request
//as an argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)

	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil

}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &Repository{}

	//Create a new service.
	srv := micro.NewService(
		micro.Name("consignment"),
	)

	srv.Init()

	//Register our service with the gRPC server, this will tie
	//our implementation into the auto-generated interface
	//for our protobuf definition
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
