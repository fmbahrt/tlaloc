package main

import (
    "flag"
    "net"
    "fmt"
    "log"

    //"golang.org/x/net/context" //old context check when grpc is updated
    "google.golang.org/grpc"
    "./escache"
    rpc "./rpcserver"
    pb "../api"
)


func main() {
    port := flag.Int("p", 8181, "port to listen to")
    flag.Parse()

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil{
        log.Fatalf("Could not listen to port %d: %v", *port, err)
    }

    log.Println("Creating new gRPC server...")
    serv := grpc.NewServer()

    log.Println("Creating Elastic Search cache...")
    escache, err := escache.NewEsCache("http://localhost:9200", "features")
    if err != nil{
        log.Fatalf("Could not create EsCache: %v", err)
    }

    rpcserver := rpc.Server{
        Es: *escache,
    }

    pb.RegisterDistanceServer(serv, rpcserver)
    log.Println("Ready...")
    err = serv.Serve(lis)
    if err != nil{
        log.Fatalf("Could not serve: %v", err)
    }
}
