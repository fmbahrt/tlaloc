package main

import (
    "log"
    "net"
    "fmt"

    "google.golang.org/grpc"
    "./escache"
    rpc "./rpcserver"
    pb "../api"
)

type Service struct {
    Port       int
    Rpcserver  rpc.Server
}

//FIXME
func NewService(port int) (*Service, error) {
    escache, err := escache.NewEsCache("http://localhost:9200", "features")

    if err != nil{
        return nil, err
    }

    rpcserver := rpc.Server{
        Es: *escache,
    }

    service := &Service{
        Port: port,
        Rpcserver: rpcserver,
    }
    return service, nil
}

// FIXME
func (s *Service) Deploy() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
    if err != nil{
        log.Fatalf("Could not listen to port %d: %v", s.Port, err)
    }

    log.Println("Creating new gRPC server...")
    serv := grpc.NewServer()

    pb.RegisterDistanceServer(serv, s.Rpcserver)
    log.Println("Ready...")
    err = serv.Serve(lis)
    if err != nil{
        log.Fatalf("Could not serve: %v", err)
    }
}
