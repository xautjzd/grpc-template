package main

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "github.com/xautjzd/grpc-template/structure"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (s *server) AddPerson(ctx context.Context, req *pb.PersonCreateRequest) (*pb.PersonCreateResponse, error) {
	logrus.Infof("Received: %v", req)
	return &pb.PersonCreateResponse{Msg: "success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
