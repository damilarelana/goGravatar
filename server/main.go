package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/damilarelana/goGravatar/gravatar"
	"github.com/damilarelana/goGravatar/hasher"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// set the Port that the server would be listening/serving from
const port = ":9091"

// GravatarService to run the service [from gravatar.proto]
type GravatarService struct {
	// ...
}

// Generate service handler method
func (s *GravatarService) Generate(ctx context.Context, in *pb.GravatarRequest) (*pb.GravatarResponse, error) {
	log.Printf("Received email %v with size %v", in.Email, in.Size)
	return &pb.GravatarResponse{Url: hasher.Gravatar(in.Email, in.Size)}, nil
}

func main() {

	// start a server listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Failed to listen on port %v", port)))
	}

	// instanstiate a gRPC server
	server := grpc.NewServer()
	pb.RegisterGravatarServiceServer(server, &GravatarService{})

	// start the gRPC Server
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start gRPC Server"))
	}
}
