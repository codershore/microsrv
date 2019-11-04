package main

import (
	pb "github.com/codershore/microsrv/vessel-service/proto/vessel"
	"os"
)

const defaultHost  = "localhost:27071"

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id:"vessel001", Name:"Kane's Salty Secret", MaxWeight:200000, Capacity:500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}


func main()  {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
}
