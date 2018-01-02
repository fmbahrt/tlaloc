package escache

import (
    "../../distance_metrics"
)

type SearchEntry struct {
	Vector []float32 `json:"vector"`
	ID     int       `json:"id"`
	Dist   float32   `json:"dist,omitempty"`
}

// TODO strategy pattern? Use higher order function to inject distmetric
func (se *SearchEntry) Distance(sc SearchEntry) (float32, error) {
    return distance_metrics.ManhattanDistance(se.Vector,
                                              sc.Vector)
}

func (se *SearchEntry) SetDist(distance float32) {
    se.Dist = distance
}

// Implement sort interface for sorting TODO TEST THIS IMPORTANT!
type ByID []SearchEntry

func (a ByID) Len() int { return len(a) }
func (a ByID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }
