package widgetserver

import (
	"context"
	"log"
	"net"

	pb "grpclib/widget"

	"google.golang.org/grpc"
)

const (
	port = ":9890"
)

type server struct {
	pb.UnimplementedProcessServer
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// Do something with the request, send to database, cache, etc...

	response := "ok"
	return &pb.UpdateResponse{Message: response}, nil
}

// Code to be placed in the storage service to listen for requests to update
func main() {
	// Listen for grpc client requests
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProcessServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
