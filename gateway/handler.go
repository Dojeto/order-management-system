package main

import (
	"log"

	"github.com/gin-gonic/gin"

	pb "github.com/dojeto/order-management-system/common/api"
)

type Handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *Handler {
	return &Handler{
		client,
	}
}

func (h Handler) CreateOrder(ctx *gin.Context) {
	body := &pb.CreateOrderRequest{}
	err := ctx.Bind(body)
	if err != nil {
		log.Fatal(err)
	}
	h.client.CreateOrder(ctx, body)
}
