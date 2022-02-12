package main

import (
	"context"
	"flag"
	pb "github.com/cdecoux/golang-backpack/protobuf/demo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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

	// Dial to gRPC Server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRandomNumberClient(conn)

	// Get Single Random Number
	number, err := client.GetRandomNumber(context.Background(), &pb.GetRandomNumberRequest{})
	if err != nil {
		return
	}
	log.Info(number)

	// Setup a Stream of Random Numbers
	requestSize := int64(1000)
	stream, err := client.GetRandomNumberStream(context.Background(), &pb.GetRandomNumberStreamRequest{Count: requestSize})
	if err != nil {
		return
	}

	for {
		randomNumber, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetRandomNumberStream(_) = _, %v", client, err)
		}

		log.Info(randomNumber)
	}

}
