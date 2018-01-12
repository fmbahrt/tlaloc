package main

import (
    "log"
    "net"
    "net/http"
    "./routing"

    "google.golang.org/grpc"

    pb "../registry"
)

func main() {

    log.Printf("Starting Service Registry")
    go func(){
        port := ":8081"
        lis, err := net.Listen("tcp", port)
        if err != nil {
            log.Fatalf("Could not listen to port %d: %v", port, err)
        }

        serv := grpc.NewServer()
        pb.RegisterRegistryServer(serv, pb.NewServiceRegistry(uint32(1)))
        err = serv.Serve(lis)
        if err != nil {
            log.Fatalf("Could not serve: %v", err)
        }
        log.Printf("Done!")
    }()

    log.Printf("Starting service...")
    router := routing.NewRouter(routing.ServerRoutes)

    log.Fatal(http.ListenAndServe(":8080", router))
}
