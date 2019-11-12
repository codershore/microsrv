package main

import (
	"context"
	pbUser "github.com/codershore/microsrv/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config/cmd"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
)
func main() {
	cmd.Init()

	client := pbUser.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:        "name",
				Usage:       "Your full name",
			},
			cli.StringFlag{
				Name:        "email",
				Usage:       "Your email",
			},
			cli.StringFlag{
				Name:        "password",
				Usage:       "Your password",
			},
			cli.StringFlag{
				Name:        "company",
				Usage:       "Your company",
			},
			),
			)

	service.Init(
		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.TODO(), &pbUser.User{
				Name:                 name,
				Company:              company,
				Email:                email,
				Password:             password,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pbUser.Request{})
			if err != nil {
				log.Fatalf("Cloud not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}
			os.Exit(0)
		}),
	)
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
