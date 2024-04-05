package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dranonymousnet/eventhorizon/api/v1/notifier"
	"github.com/dranonymousnet/eventhorizon/internal/config"
	"github.com/dranonymousnet/eventhorizon/internal/store"
)

type NotifierServer struct {
	notifier.UnimplementedNotifierServiceServer

}

func (NotifierServer) 	Notify(context.Context, *notifier.NotifyRequest) (*notifier.NotifyResponse, error){
	return &notifier.NotifyResponse{}, nil
}

func init(){
	config.Setup()
	store.InitRedis()
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen on port")
	}

	serverRegistrar := grpc.NewServer()
	service := &NotifierServer{}
	notifier.RegisterNotifierServiceServer(
		serverRegistrar, service,
	)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve")
	}
	log.Println("Server started")

}
