package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sdual/mlserving/apps/serving/adaptor/controller"
	pb "github.com/sdual/mlserving/proto/grpc/serving/predict"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
}

func (gs GRPCServer) Start() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterPredictServiceServer(s, &controller.PredictServiceServer{})
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
