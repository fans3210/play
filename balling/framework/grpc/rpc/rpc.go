package rpc

import (
	"balling/balling/framework/grpc/pb"
	dm "balling/domain"
	"balling/domain/constants"
	"balling/domain/errs"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func gameFromPbRequest(req *pb.CalculateScoreRequest) (dm.Game, error) {

	ipt := make([][]uint32, 0)
	fs := req.GetGame().GetFrames()
	for _, f := range fs {
		tr := f.GetThrows()
		ipt = append(ipt, tr)
	}

	if len(ipt) != constants.NumFrames {
		return dm.Game{}, errs.ErrInvalidInput
	}

	var ret dm.Game
	copy(ret[:], ipt)

	return ret, nil
}

func (s *server) CalculateScore(ctx context.Context, req *pb.CalculateScoreRequest) (*pb.CalculateScoreResponse, error) {

	// validate input
	ipt, err := gameFromPbRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

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
