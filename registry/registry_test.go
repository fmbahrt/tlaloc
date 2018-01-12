package registry

import (
    "context"
    "testing"
)

var registerTest = []struct{
    name            string
    endpoints       []Endpoint
    serviceRegistry *ServiceRegistry
    ctx             context.Context
}{
    {"Add one endpoint",
     []Endpoint{
         Endpoint{
            Address: "localhost",
            Port: int32(80),
         },
     },
     NewServiceRegistry(uint32(1)),
     context.Background()},
    {"Add two endpoints",
     []Endpoint{
         Endpoint{
            Address: "localhost",
            Port: int32(80),
         },
         Endpoint{
            Address: "localhooost",
            Port: int32(100),
         },
     },
     NewServiceRegistry(uint32(1)),
     context.Background()},
    {"Add multiple endpoints",
     []Endpoint{
         Endpoint{
            Address: "localhost",
            Port: int32(80),
         },
         Endpoint{
            Address: "localhooost",
            Port: int32(100),
         },
         Endpoint{
            Address: "search_slave",
            Port: int32(101),
         },
         Endpoint{
            Address: "192.168.1.2",
            Port: int32(4000),
         },
     },
     NewServiceRegistry(uint32(1)),
     context.Background()},
}

func TestRegister(t *testing.T){
    for _, tt := range registerTest{
        t.Run(tt.name, func(t *testing.T){
            // Act
            for i, end := range tt.endpoints{
                lease,_ := tt.serviceRegistry.Register(tt.ctx, &end)

                // Assert one services length right now
                // Test for correct lease aswell
                if lease.CheckInInterval != tt.serviceRegistry.CheckInInterval {
                    t.Errorf("Lease CheckInInteval != ServiceRegistry CheckInInterval")
                }

                if len(tt.serviceRegistry.services) != (i+1) {
                    t.Fatalf("Service not appened to services slice")
                }
            }
        })
    }
}

var unregisterTest = []struct{
    name            string
    lease           *Lease
    serviceRegistry *ServiceRegistry
    services        []*Service
    ctx             context.Context
    confExpected    bool
}{
    {"Remove a lease when total number of services is 1",
     &Lease{
        Id: "1337",
        CheckInInterval: 1,
     },
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
     },
     context.Background(),
     true,
     },
    {"Remove a lease when total number of services is 1, lease does not exists",
     &Lease{
        Id: "4000",
        CheckInInterval: 1,
     },
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
     },
     context.Background(),
     false,
     },
    {"Remove a lease when total number of services is 2",
     &Lease{
        Id: "1337",
        CheckInInterval: 1,
     },
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
         &Service{
            Id: "3000",
         },
     },
     context.Background(),
     true,
     },
    {"Remove a lease when total number of services is 2, lease does not exists",
     &Lease{
        Id: "4000",
        CheckInInterval: 1,
     },
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
         &Service{
            Id: "3000",
         },
     },
     context.Background(),
     false,
     },
}

func TestUnregister(t *testing.T){
    for _, tt := range unregisterTest{
        t.Run(tt.name, func(t *testing.T){
            // Assert on slice length right now
            // Change this.

            // Arrange
            tt.serviceRegistry.services = tt.services

            // Act
            conf,_ := tt.serviceRegistry.Unregister(tt.ctx, tt.lease)

            // Assert
            if !(conf.Ok == tt.confExpected) {
                t.Fatalf("Failed to remove service")
            }
        })
    }
}

var getAllServicesTest = []struct{
    name            string
    serviceRegistry *ServiceRegistry
    services        []*Service
    ctx             context.Context
    elemsExpected   int
}{
    {"Test with one elements in registry",
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
     },
     context.Background(),
     1,
     },
    {"Test with two elements in registry",
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
         &Service{
            Id: "213123",
         },
     },
     context.Background(),
     2,
     },
    {"Test with multiple elements in registry",
     NewServiceRegistry(uint32(1)),
     []*Service{
         &Service{
            Id: "1337",
         },
         &Service{
            Id: "213123",
         },
         &Service{
            Id: "133asd7",
         },
         &Service{
            Id: "213123k12j3",
         },
     },
     context.Background(),
     4,
     },
    {"Test with one elements in registry",
     NewServiceRegistry(uint32(1)),
     []*Service{},
     context.Background(),
     0,
     },
}

func TestGetAllServices(t *testing.T){
    for _, tt := range getAllServicesTest{
        t.Run(tt.name, func(t *testing.T){
            // Arrange
            tt.serviceRegistry.services = tt.services

            services,_ := tt.serviceRegistry.GetAllServices(tt.ctx, &EmptyParam{})

            if len(services.Services) != tt.elemsExpected {
                t.Fatalf("Len of service not as expected")
            }
        })
    }
}
