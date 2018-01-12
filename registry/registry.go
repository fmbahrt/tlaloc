package registry

import (
    "sync"
    "context"

    "github.com/chilts/sid"
)

//type Service struct {
//    Id       string
//    Endpoint Endpoint
//}

type ServiceRegistry struct {
    services        []*Service
    CheckInInterval uint32
    sync.RWMutex // read-write mutex
}

func NewServiceRegistry(checkInInterval uint32) (*ServiceRegistry) {
    // What if checkInInterval is zero?
    // checkInInterval is in seconds
    sr := &ServiceRegistry{
        CheckInInterval: checkInInterval,
    }
    return sr
}

func (s *ServiceRegistry) Register(ctx context.Context, endpoint *Endpoint) (*Lease, error) {
    id := sid.Id()

    service := &Service{
        Id: id,
        Endpoint: endpoint,
    }

    lease := &Lease{
        Id: id,
        CheckInInterval: s.CheckInInterval,
    }

    s.Lock()
    s.services = append(s.services, service)
    defer s.Unlock()

    return lease, nil
}

func (s *ServiceRegistry) Unregister(ctx context.Context, lease *Lease) (*Confirmation, error) {
    // Find service by id then remove it from services
    removed := false

    s.Lock()
    defer s.Unlock()

    for i, service := range s.services {
        if service.Id == lease.Id {
            s.services = append(s.services[:i], s.services[i+1:]...)
            removed = true
            break; // Perhaps use while loop instead
        }
    }

    // FIXME I do not like the Confirmation struct - find alternative
    confirmation := &Confirmation{
        Ok: removed,
    }

    return confirmation, nil
}

// TODO not done
func (s *ServiceRegistry) GetAllServices(ctx context.Context, ep *EmptyParam) (*Services, error) {
    services := &Services{
        Services: s.services,
    }
    return services, nil
}

// This has not been implemented
func (s *ServiceRegistry) CheckIn(ctx context.Context, lease *Lease) (*Confirmation, error) {
    return &Confirmation{Ok: true}, nil
}

func (s *ServiceRegistry) TotalServices() (int) {
    s.RLock()
    defer s.RUnlock()
    return len(s.services)
}
