package main

import (
	"context"
	pbUser "github.com/codershore/microsrv/user-service/proto/user"
)
type service struct {
	repo Reposityory
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req pbUser.User, res *pbUser.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, request pbUser.Request, res *pbUser.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return  nil
}

func (srv *service) Auth(ctx context.Context, req *pbUser.User, res *pbUser.Token) error {
	_, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}

func (srv *service) Create(ctx context.Context, req *pbUser.User, response pbUser.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	response.User = req
	return nil
}