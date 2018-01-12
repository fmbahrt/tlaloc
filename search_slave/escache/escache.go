package escache

import (
    "io"
    "fmt"
    "time"
    "sync"
    "sort"
    "encoding/json"

    "golang.org/x/sync/errgroup"
    "golang.org/x/net/context"

    "github.com/olivere/elastic"
)

type EsCache struct {
    sync.RWMutex
    LastUpdate  time.Time
    Entries     []SearchEntry //Keep sorted by id in ascending order
    EsClient    *elastic.Client
    EsIndex     string
}

func NewEsCache(esurl string, index string) (*EsCache, error) {
    //http://localhost:9200/
	client, err := elastic.NewClient(elastic.SetURL(esurl),
                                     elastic.SetSniff(false))
	if err != nil {
		// Handle error
        return nil, fmt.Errorf("ES client error: %v", err)
	}

    escache := &EsCache{
        LastUpdate: time.Now(),
        EsClient: client,
        EsIndex: index,
    }

    err = escache.Update()
    if err != nil {
        return nil, err
    }

    return escache, nil
}

// Find SearchEntry by performing binary search on id
func (es *EsCache) GetEntryById(id int) (*SearchEntry, error) {
    //Assume - as of right now - this cache have all data
    i := sort.Search(len(es.Entries), func(i int) bool {
        return es.Entries[i].ID >= id
    })
    if i == len(es.Entries){
        //Id no founds
        return nil, fmt.Errorf("Id does not exist")
    } else{
        return &es.Entries[i], nil
    }
}

func (es *EsCache) Update() error{
    /*
    ctx := context.Background()

	sr, err := es.EsClient.Search().
		Index(es.EsIndex).   // search in index "features"
        From(1).Size(9999).
		Do(ctx)             // execute

	if err != nil {
		// Handle error
		return fmt.Errorf("ES search result error: %v\n", err)
	}

    if sr.Hits.TotalHits > 0 {

        // Iterate through results
        for _, hit := range sr.Hits.Hits {
            // hit.Index contains the name of the index

            se := SearchEntry{}
            err := json.Unmarshal(*hit.Source, &se)
            if err != nil {
                // Deserialization failed
                return fmt.Errorf("Deserialization failed: %v\n", err)
            } else { // No errors
                es.Entries = append(es.Entries,se)
            }
        }
    }*/

    hits, eg, ctx := es.hitsToChannel(context.Background(), *es.EsClient)

    for i := 0; i < 10; i++{
        es.deserializeHits(ctx, eg, hits)
    }

    // check if any go-routine failed
    if err := eg.Wait(); err != nil {
        panic(err) // Log this and retry?
    }

    //Sorted insert might be better
    sort.Sort(ByID(es.Entries))
    es.LastUpdate = time.Now()
    return nil
}

func (es *EsCache) hitsToChannel(ctx context.Context, client elastic.Client) (<-chan json.RawMessage, *errgroup.Group, context.Context) {
    out := make(chan json.RawMessage)

    eg, ctx := errgroup.WithContext(ctx) // Place this outside function?

    eg.Go(func() error {
        defer close(out)
        // Initialize scroller. Just don't call Do yet.
        scroll := client.Scroll("features").Size(100) // Make this input args - dont hardcode
        for {
            results, err := scroll.Do(ctx)
            if err == io.EOF { // end of file
                return nil // all results retrieved
            }
            if err != nil {
                return err // something went wrong
            }

            // Send the hits to the hits channel
            for _, hit := range results.Hits.Hits {
                select {
                    case out <- *hit.Source:
                    case <-ctx.Done():
                    return ctx.Err()
                }
            }
        }
    })

    return out, eg, ctx
}

func (es *EsCache) deserializeHits(ctx context.Context, eg *errgroup.Group, in <-chan json.RawMessage) {
	eg.Go(func() error {
		for hit := range in {
            se := SearchEntry{}
            err := json.Unmarshal(hit, &se)
            if err != nil {
                // Deserialization failed
                return err
            }
			es.Lock()
			es.Entries = append(es.Entries,se)
            es.Unlock()

            // Terminate early is error occurs somewhere in the pipeline
            select {
                default:
                case <-ctx.Done():
                    return ctx.Err()
            }
		}
        return nil
	})
}

func (es *EsCache) DataAsChannel(wg *sync.WaitGroup) <-chan SearchEntry{
    ch := make(chan SearchEntry, 20) //Maybe buffered?

    go func(){
        for _, elem := range es.Entries{
            ch <- elem
        }
        close(ch)
        wg.Done()
    }()

    return ch
}
