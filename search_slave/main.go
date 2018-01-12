package main

import (
    "log"
    "flag"
    "time"
    "context"
    //"golang.org/x/net/context" //old context check when grpc is updated

    "google.golang.org/grpc"

    reg "../registry"
)


func main() {
    port := flag.Int("p", 8181, "port to listen to")
    flag.Parse()

    service, err := NewService(*port)
    if err != nil {
        log.Fatalf("Failed to create new service: %v", err)
    }

    // register in service registry
    conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
    if err != nil{
        log.Fatalf("Could not connect to service registry")
        return
    }
    defer conn.Close()

    srClient := reg.NewRegistryClient(conn)
    lease,_ := srClient.Register(context.Background(), &reg.Endpoint{
        Address: "localhost",
        Port: int32(*port),
    })

    log.Printf("Got a lease: %v", lease)

    // Sending heartbeats
    go func(){
        // FIXME How do I stop this properly?
        for {
            log.Printf("Sending heartbeat...")
            srClient.CheckIn(context.Background(), lease)
            <-time.After(time.Duration(lease.CheckInInterval) * time.Second)
        }
    }()

    service.Deploy()
}
