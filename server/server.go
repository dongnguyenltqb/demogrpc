package server

import (
	"context"
	"demogrpc/rpc"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GRPCClient struct {
	UserRPCClient rpc.UserClient
}

var GRPC_CLI GRPCClient

func createGRPCClient() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	GRPC_CLI = GRPCClient{
		UserRPCClient: rpc.NewUserClient(conn),
	}
}

func Start() {
	createGRPCClient()
	cli := GRPC_CLI.UserRPCClient
	app := gin.Default()
	u := app.Group("/")
	{
		u.POST("/login", func(c *gin.Context) {
			var cre Credentials
			if err := c.BindJSON(&cre); err != nil {
				fmt.Println(err)
				c.JSON(400, gin.H{
					"Ok":      false,
					"Message": err.Error(),
				})
				return
			}
			response, err := cli.UserLogin(context.Background(), &rpc.Credentials{
				Username: cre.Email,
				Password: cre.Password,
			})
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{
					"ok":      false,
					"message": err.Error,
				})
				return
			}
			c.JSON(200, response)
		})
	}
	app.Run(":7000")
}
