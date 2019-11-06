package main

import (
	pbUser "github.com/codershore/microsrv/user-service/proto/user"
)
type service struct {
	repo Repository
	tokenService Authable
}
