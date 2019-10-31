package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/codershore/microsrv/consignment-service/proto/consignment"
	vesselProto "github.com/codershore/microsrv/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

//const (
//	port = ":50051"
//)

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
	vesselCient vesselProto.VesselServiceClient
}

//CreateConsignment - we created just on method on our service,
//Which is a create method, which takes a context and a request
//as an argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselCient.FindAvailable(context.Background(), &vesselProto.Specification{
		Capacity:             int32(len(req.Container)),
		MaxWeight:            req.Weight,
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

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
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}


}
