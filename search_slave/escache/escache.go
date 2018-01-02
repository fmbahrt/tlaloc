package escache

import (
    "fmt"
    "time"
    "context"
    "sync"
    "sort"
    "encoding/json"

    "github.com/olivere/elastic"
)

type EsCache struct { //Perharps lock this
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
    }
    //Sorted insert might be better
    sort.Sort(ByID(es.Entries))
    es.LastUpdate = time.Now()
    return nil
}

func (es *EsCache) DataAsChannel(wg *sync.WaitGroup) <-chan SearchEntry{
    ch := make(chan SearchEntry) //Maybe buffered?

    go func(){
        for _, elem := range es.Entries{
            ch <- elem
        }
        close(ch)
        wg.Done()
    }()

    return ch
}
