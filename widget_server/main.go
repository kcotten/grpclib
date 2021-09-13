package widgetserver

import (
	pb "github.com/kcotten/grpclib/widget"
)

const (
	port = ":9890"
)

type server struct {
	pb.UnimplementedProcessServer
}
