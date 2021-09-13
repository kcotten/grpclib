package widgetclient

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpclib/widget"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9890"
)

// Code for the rest endpoint to send request to the storage service
// Alternately all the code could be in a monolithic package
// 1. Unmarshall to protobuf
// 2. Send request
// 3. Marshall response
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProcessClient(conn)

	// Contact the grpc server
	name := "item1"
	quantity := 7
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Update(ctx, &pb.UpdateRequest{Name: name, Quantity: int64(quantity)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
