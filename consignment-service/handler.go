package main

import (
	"context"
	pb "github.com/codershore/microsrv/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
	"log"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository  {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		log.Fatalln("Create consignment error %v", err)
		return err
	}
	res.Consignment = req
	res.Created = true
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, response *pb.Response) error{
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	response.Consignments = consignments
	return nil
}
