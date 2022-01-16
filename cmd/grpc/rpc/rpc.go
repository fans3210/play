package rpc

import (
	"balling/balling/framework/grpc/pb"
	dm "balling/domain"
	"balling/domain/errs"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func gameFromPbRequest(req *pb.CalculateScoreRequest) dm.Game {

	ret := dm.Game{}
	fs := req.GetGame().GetFrames()
	for i, f := range fs {
		tr := f.GetThrows()
		ret[i] = tr
	}
	return ret
}

func (s *server) CalculateScore(ctx context.Context, req *pb.CalculateScoreRequest) (*pb.CalculateScoreResponse, error) {

	ipt := gameFromPbRequest(req)
	uc := s.f.MakeCalculateBallingScoreUseCase(ipt)
	res, err := uc.Run()
	if err != nil {
		switch err {
		case errs.ErrInvalidInput:
			errMsg := fmt.Sprintf("invalid input : %v", ipt)
			e := status.Error(codes.InvalidArgument, errMsg)
			return nil, e
		case errs.ErrUnexpected:
			errMsg := fmt.Sprintf("unexpected error: %v, input: %v", err, ipt)
			e := status.Error(codes.Unknown, errMsg)
			return nil, e
		default:
			errMsg := fmt.Sprintf("unhandled error: %v, input: %v", err, ipt)
			e := status.Error(codes.Internal, errMsg)
			return nil, e
		}
	}

	return &pb.CalculateScoreResponse{
		Results: res[:],
	}, nil
}
