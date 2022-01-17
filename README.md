Project Structutes:

├── README.md
├── balling.features
├── balling
│   ├── app
│   │   ├── core.go
│   │   └── core_test.go
│   ├── framework
│   │   └── grpc
│   │       ├── pb
│   │       │   ├── balling.pb.go
│   │       │   ├── ballingService.pb.go
│   │       │   └── ballingService_grpc.pb.go
│   │       └── proto
│   │           ├── balling.proto
│   │           └── ballingService.proto
│   └── usecase
│       ├── cal_ballingscore.go
│       └── cal_ballingscore_test.go
├── cmd
│   ├── cli
│   │   └── main.go
│   └── grpc
│       ├── main.go
│       └── rpc
│           ├── rpc.go
│           ├── rpc_e2e_test.go
│           └── server.go
├── di
│   ├── prod_container.go
│   └── test_container.go
├── domain
│   ├── balling.go
│   ├── constants
│   │   └── constants.go
│   └── errs
│       └── errors.go
├── go.mod
├── go.sum
└── pb_gen.sh




