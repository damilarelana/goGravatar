package main

import (
	"context"
	"log"
	"time"

	pb "github.com/damilarelana/goGravatar/gravatar"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// set the address that the client would be communicating through
const address = "localhost:9091"

func main() {

	// instantiate and connect a gRPC Client to connect the grPC Server address
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to connect via grpc"))
	}
	defer conn.Close() // defer the close of the connection until main() is terminated

	//
	c := pb.NewGravatarServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Generate(ctx, &pb.GravatarRequest{Email: "kamil@lelonek.me", Size: 10})
	if err != nil {
		log.Fatal(errors.Wrap(err, "Could not greet ... ")
	}

	log.Printf("Greeting: %s", r.Url)
}
