package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dojeto/order-management-system/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	config := Config{
		URL: "mongodb://mongodb:27017/orders",
	}
	store := NewStore(config)
	store.Create(context.TODO())
	fmt.Println(err)
	// select {}
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
