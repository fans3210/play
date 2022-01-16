package main

import (
	"balling/cmd/grpc/rpc"
	"balling/di"
)

func main() {
	c := di.NewProductionContainer()
	s := rpc.NewGRPCServer(c)
	rpc.Start(s)
}
