package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/cdecoux/golang-backpack/protobuf/demo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"net"
	"os"
)

var port = flag.Int("port", 50051, "The server port")

func main() {

	/*
		Set the Log Level of LogRUS
	*/
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "INFO"
	}

	logLevel, err := log.ParseLevel(lvl)
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)

	log.Info("Creating the GRPC Server")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRandomNumberServer(grpcServer, &demoServer{})
	grpcServer.Serve(lis)
	log.Info("Finished Starting Server")

}

type demoServer struct {
	pb.UnimplementedRandomNumberServer
}

func (server *demoServer) GetRandomNumber(ctx context.Context, req *pb.GetRandomNumberRequest) (*pb.GetRandomNumberResponse, error) {

	res := &pb.GetRandomNumberResponse{
		RandomNumber: rand.Int63(),
		Time:         timestamppb.Now(),
	}
	return res, nil
}

func (server *demoServer) GetRandomNumberStream(req *pb.GetRandomNumberRequest, stream pb.RandomNumber_GetRandomNumberStreamServer) (err error) {
	return status.Errorf(codes.Unimplemented, "method GetRandomNumberStream not implemented")
}
