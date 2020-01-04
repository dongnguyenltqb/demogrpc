package server

import (
	"context"
	"fmt"
	"log"
	"share-proto/proto-gen/message"
	"share-proto/proto-gen/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FormRegister struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" bindding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type GRPCClient struct {
	UserRPCClient rpc.UserClient
}

var GRPC_CLI *GRPCClient

func createGRPCClient() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	GRPC_CLI = &GRPCClient{
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
			response, err := cli.UserLogin(context.Background(), &message.Credentials{
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
		u.POST("/register", func(c *gin.Context) {
			var info FormRegister
			if err := c.BindJSON(&info); err != nil {
				fmt.Println(err)
				c.JSON(400, gin.H{
					"Ok":      false,
					"Message": err.Error(),
				})
				return
			}
			response, err := cli.UserRegister(context.Background(), &message.FormRegister{
				FirstName:   info.FirstName,
				LastName:    info.LastName,
				Email:       info.Email,
				Password:    info.Password,
				PhoneNumber: info.PhoneNumber,
				Address:     info.Address,
			})
			if err != nil {
				c.JSON(500, gin.H{
					"Ok":      false,
					"Message": err.Error(),
				})
				return
			}
			c.JSON(200, response)
		})
	}
	app.Run(":7000")
}
