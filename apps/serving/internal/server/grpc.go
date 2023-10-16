package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sdual/mlserving/apps/serving/internal/adaptor/controller"
	"github.com/sdual/mlserving/apps/serving/internal/config"
	pb "github.com/sdual/mlserving/proto/grpc/serving/predict"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
}

func (gs GRPCServer) Start(config config.GRPCConfig) {
	port := config.Port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterPredictServiceServer(server, &controller.PredictServiceServer{})
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %d", port)
		server.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
