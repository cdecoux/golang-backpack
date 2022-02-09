package main

import (
	"context"
	"flag"
	pb "github.com/cdecoux/golang-backpack/protobuf/demo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

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

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRandomNumberClient(conn)

	number, err := client.GetRandomNumber(context.Background(), &pb.GetRandomNumberRequest{})
	if err != nil {
		return
	}

	log.Info(number)
}
