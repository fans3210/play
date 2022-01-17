package rpc_test

import (
	"balling/balling/framework/grpc/pb"
	"balling/balling/framework/grpc/rpc"
	"balling/di"
	dm "balling/domain"
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
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

func pbRequestParamFromInput(ipt dm.Game) pb.CalculateScoreRequest {
	frames := make([]*pb.Frame, len(ipt))
	for i, f := range ipt {
		pbf := &pb.Frame{
			Throws: f,
		}
		frames[i] = pbf
	}

	return pb.CalculateScoreRequest{
		Game: &pb.Game{
			Frames: frames,
		},
	}

}

func TestValidation(t *testing.T) {

	ctx := context.Background()
	conn := getGRPCConn(ctx, t)
	defer conn.Close()
	client := pb.NewBallingServiceClient(conn)

	input := dm.Game{{5, 2}, {8, 1}, {6, 4}, {10}, {0, 5}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {10, 2, 6}}
	assert := assert.New(t)
	var ipt dm.Game
	var err error

	ipt = input
	params := pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.NoError(err)

	// frame 1-10, num throws = 0
	ipt = input
	ipt[2] = make([]uint32, 0)
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 1-9, num throws >=3,
	ipt = input
	ipt[1] = []uint32{1, 2, 3}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 10 num throws >= 4
	ipt = input
	ipt[9] = []uint32{1, 2, 3, 4}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 1-9, num throws = 1 but 1st throw is not strike
	ipt = input
	ipt[2] = []uint32{5}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 10, num throws = 3 but 1st throw is not strike
	ipt = input
	ipt[9] = []uint32{1, 2, 3}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 10, num throws < 3 but first throw is strike
	// The last frame has three throws only if a bowler makes a strike on the first throw. means 10, 10, 10 is valid
	ipt = input
	ipt[9] = []uint32{10, 9}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 1-9, score sum > 10(maximum)
	ipt = input
	ipt[1] = []uint32{6, 9}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// frame 10, score sum > 30(maximum)
	ipt = input
	ipt[9] = []uint32{10, 10, 11}
	params = pbRequestParamFromInput(ipt)
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// incomplete frame: less than 10 frames,
	incompleteIpt := [][]uint32{{1}}
	frames := make([]*pb.Frame, len(incompleteIpt))
	for i, f := range incompleteIpt {
		pbf := &pb.Frame{
			Throws: f,
		}
		frames[i] = pbf
	}
	params = pb.CalculateScoreRequest{
		Game: &pb.Game{
			Frames: frames,
		},
	}
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

	// more than 10 frames,
	overCompleteIpt := input[:]
	overCompleteIpt = append(overCompleteIpt, []uint32{1, 2})
	frames = make([]*pb.Frame, len(overCompleteIpt))
	for i, f := range overCompleteIpt {
		pbf := &pb.Frame{
			Throws: f,
		}
		frames[i] = pbf
	}
	params = pb.CalculateScoreRequest{
		Game: &pb.Game{
			Frames: frames,
		},
	}
	_, err = client.CalculateScore(ctx, &params)
	assert.Error(err)

}

func TestCalculate(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConn(ctx, t)
	defer conn.Close()
	client := pb.NewBallingServiceClient(conn)

	ipt := dm.Game{{5, 2}, {8, 1}, {6, 4}, {0, 0}, {0, 5}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {2, 6}}
	params := pbRequestParamFromInput(ipt)
	res, err := client.CalculateScore(ctx, &params)
	r := res.GetResults()
	assert.NoError(t, err)
	expected := dm.Scores{7, 16, 26, 26, 31, 39, 48, 56, 63, 71}
	assert.ElementsMatch(t, r, expected)

	ipt = dm.Game{{5, 2}, {8, 1}, {6, 4}, {10}, {0, 5}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {10, 2, 6}}
	params = pbRequestParamFromInput(ipt)
	res, err = client.CalculateScore(ctx, &params)
	r = res.GetResults()
	assert.NoError(t, err)
	expected = dm.Scores{7, 16, 26, 41, 46, 54, 63, 71, 78, 96}
	assert.ElementsMatch(t, r, expected)

	// consecutive strike case
	ipt = dm.Game{{5, 2}, {8, 1}, {6, 4}, {10}, {10}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {10, 2, 6}}
	params = pbRequestParamFromInput(ipt)
	res, err = client.CalculateScore(ctx, &params)
	r = res.GetResults()
	assert.NoError(t, err)
	expected = dm.Scores{7, 16, 26, 54, 72, 80, 89, 97, 104, 122}
	assert.ElementsMatch(t, r, expected)

}
