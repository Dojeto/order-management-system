package main

import (
	"context"
	"log"

	pb "github.com/dojeto/order-management-system/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order Received ! Order %v", p)
	o := &pb.Order{
		ID: "12",
		Items: []*pb.Item{
			&pb.Item{
				ID:       "12",
				PriceID:  "23",
				Quantity: 2,
				Name:     "ndjsads",
			},
			&pb.Item{
				ID:       "12",
				PriceID:  "23",
				Quantity: 2,
				Name:     "ndjsads",
			},
		},
	}
	return o, nil
}
