package main

import (
	"context"
	vesselProto "github.com/codershore/microsrv/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository  {
	return &VesselRepository{s.session.Clone()}

}
func (s *service) FindAvailable(ctx context.Context, req *vesselProto.Specification, res *vesselProto.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return  err
	}

	res.Vessel = vessel
	return nil

}

func (s *service) Create(ctx context.Context, req *vesselProto.Vessel, res *vesselProto.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}