package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
)

type Server struct {
}

func (s *Server) InfoObtained(ctx context.Context, message *Info) (*Info, error) {
	log.Printf("Received stundent information: %s", message.Body)
	return &Info{Body: "Hi, we have received your information"}, nil

}
