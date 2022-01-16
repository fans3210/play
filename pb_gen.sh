#!/bin/bash

protoc \
    -I ./balling/framework/grpc/proto \
    --go_out=./balling/framework/grpc/ \
    --go-grpc_out=./balling/framework/grpc/ \
    ./balling/framework/grpc/proto/*
