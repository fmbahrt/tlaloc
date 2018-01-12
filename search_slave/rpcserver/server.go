package rpcserver

import (
    "log"
    "context"
    "sync"
    "sort"
    "strconv"
    "../escache"

    //"golang.org/x/net/context" //old context check when grpc is updated
    pb "../../api"
)

var wg = sync.WaitGroup{}

type Server struct{
    Es  escache.EsCache
}

func (s Server) Dist(ctx context.Context, q *pb.Query) (*pb.Response, error){

    n := q.Top

    id, err := strconv.Atoi(q.Id)
    if err != nil{
        log.Printf("Cannot convery integer to string: %v", err)
        return nil, err
    }

    se, err := s.Es.GetEntryById(id)
    if err != nil{
        log.Printf("Cannot get entry by id: %v", err)
        return nil, err
    }

    wg.Add(1)
    c := s.Es.DataAsChannel(&wg)

    wg.Add(8)
    c1 := CalcDistances(*se, c)
    c2 := CalcDistances(*se, c)
    c3 := CalcDistances(*se, c)
    c4 := CalcDistances(*se, c)
    c5 := CalcDistances(*se, c)
    c6 := CalcDistances(*se, c)
    c7 := CalcDistances(*se, c)
    c8 := CalcDistances(*se, c)

    var dataset []*pb.ResponseElement
    response := &pb.Response{
        Responses: dataset,
    }

	for elem := range merge(c1,c2,c3,c4,c5,c6,c7,c8){
        re := &pb.ResponseElement{
            Id: strconv.Itoa(elem.ID),
            Dist: elem.Dist,
        }
        response.Responses = append(response.Responses, re)
	}
	wg.Wait()

    /*
        Make minheap O(n)
        Extract k elements from minheap O(k * log(n))
        Total O(n + k*log(n))
    */

    //Sort by distance
    sort.Sort(ByDistance(response.Responses))

    // FIXME: BAD - just for testing
    response.Responses = response.Responses[:n]

    //Reply
    return response, nil
}

// THIS DOES NOT BELONG HERE
type ByDistance []*pb.ResponseElement

func (a ByDistance) Len() int { return len(a) }
func (a ByDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Dist < a[j].Dist }

func merge(cs ...<-chan escache.SearchEntry) <-chan escache.SearchEntry {
    var wg sync.WaitGroup
    out := make(chan escache.SearchEntry)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan escache.SearchEntry) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

// Should be able to spawn multiple of these
func CalcDistances(se escache.SearchEntry, in <-chan escache.SearchEntry) <-chan escache.SearchEntry {
	ch := make(chan escache.SearchEntry)

	go func(){
		for elem := range in {
			dist, _ := elem.Distance(se)
			elem.Dist = dist
			ch <- elem
		}
		close(ch)
		wg.Done()
	}()
	return ch
}
