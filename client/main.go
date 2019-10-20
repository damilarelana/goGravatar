package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	mux "github.com/gorilla/mux"
)

// set the Port that the server would be listening/serving from
const port = ":9091"

// integrate the defined GravatarService [from gravatar.proto]
type GravatarService struct {
	// ...
}

// implement the service handler
func (s *GravatarService) Generate(ctx context.Context, in *pb.GravatarRequest) *pb.GravatarResponse {
	log.Printf("Received email %v with size %v", in.Email, in.Size)
	return &pb.GravatarResponse{
		Url: Gravatar(in.Email, in.Size)
	}
}

func main() {

	// start a server listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to listen on port %v due to %v", port))
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
