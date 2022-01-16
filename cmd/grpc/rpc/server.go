package rpc

import (
	"balling/balling/framework/grpc/pb"
	dm "balling/domain"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	f dm.CalBallingScoreUseCaseFactory

	pb.UnimplementedBallingServiceServer
}

func NewGRPCServer(f dm.CalBallingScoreUseCaseFactory) *server {
	return &server{
		f: f,
	}
}

func Start(srv *server) {
	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("fail to listen server, %v", err)
	}

	s := grpc.NewServer()

	// TODO: remove in production
	reflection.Register(s)

	pb.RegisterBallingServiceServer(s, srv)
	fmt.Println("server listening at: ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
