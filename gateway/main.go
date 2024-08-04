package main

import (
	"log"
	"net/http"

	"github.com/dojeto/order-management-system/common"
	pb "github.com/dojeto/order-management-system/common/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	orderServiceAddress = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	router := gin.Default()
	conn, err := grpc.NewClient(orderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewOrderServiceClient(conn)
	handler := NewHandler(client)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "I am Alive YEYE",
		})
	})
	router.POST("/", handler.CreateOrder)
	router.Run()
}
