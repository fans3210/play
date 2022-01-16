package rpc

import (
	"balling/balling/framework/grpc/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CalculateScore(ctx context.Context, req *pb.CalculateScoreRequest) (*pb.CalculateScoreResponse, error) {

	fmt.Println("hello there from grpc server")
	frames := req.GetGame().GetFrames()
	for idx, frame := range frames {
		throws := frame.GetThrows()
		fmt.Println("input = ", idx, throws)
	}
	err := status.Error(codes.Aborted, "test error")
	return nil, err
}
