package rpc_test

import (
	"balling/balling/framework/grpc/pb"
	"balling/cmd/grpc/rpc"
	"balling/di"
	"context"
	"log"
	"net"
	"testing"

	// "github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	c := di.NewTestContainer()
	srv := rpc.NewGRPCServer(c)
	pb.RegisterBallingServiceServer(s, srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDailer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConn(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDailer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dail bufnet: %v", err)
	}

	return conn
}

func TestCalculate(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConn(ctx, t)
	defer conn.Close()

	client := pb.NewBallingServiceClient(conn)
	frame1 := &pb.Frame{
		Throws: []uint32{1, 2},
	}
	params := &pb.CalculateScoreRequest{
		Game: &pb.Game{
			Frames: []*pb.Frame{frame1},
		},
	}

	res, err := client.CalculateScore(ctx, params)

	if err != nil {
		t.Fatalf("failed to calculate score : %v", err)
	}

	t.Logf("\n calculate score res is: %v", res)
}
