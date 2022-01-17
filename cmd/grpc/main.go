package main

import (
	"balling/balling/framework/grpc/rpc"
	"balling/di"
)

func main() {
	c := di.NewProductionContainer()
	s := rpc.NewGRPCServer(c)
	rpc.Start(s)
}
