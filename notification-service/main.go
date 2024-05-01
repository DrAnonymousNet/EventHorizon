package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dranonymousnet/eventhorizon/api/v1/notifier"
	"github.com/dranonymousnet/eventhorizon/internal/config"
	"github.com/dranonymousnet/eventhorizon/internal/store"
	"github.com/dranonymousnet/eventhorizon/messaging"
)

type NotifierServer struct {
	notifier.UnimplementedNotifierServiceServer
}

func (NotifierServer) Notify(ctx context.Context, requestData *notifier.NotifyRequest) (*notifier.NotifyResponse, error) {
	messaging.DispatchNotification(requestData)
	return &notifier.NotifyResponse{
		Message: "guyftdfighj",
	}, nil
}

func init() {
	config.Setup()
	store.DBSetup()
	store.InitRedis()
	messaging.Setup()
}

func main() {
	defer store.CloseDBConn()
	defer store.CloseRedisConn()


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
