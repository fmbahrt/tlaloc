package slice_utils

import (
    "container/heap"
)

type Comparable interface {
    // Compare to returns int
    // -1 if less than
    // 0 if equals
    // 1 if greater than
    CompareTo(Comparable) int
}

// This function merges k sorted slices into a single slice of length n
// The function will only return the n smallest elements among all of the
// slices.
func MergeKSortedSlicesTopN(n int, slices ...[]Comparable) []Comparable {
    // Number of slices to be merged 
    k := len(slices)

    // Check for sane input 
    // if n is larger than the combined length of the slices then
    // set n equals the combined length of the slices
    total_length := 0
    for i := 0; i < k; i++ {
        total_length = total_length + len(slices[i])
    }

    if total_length < n {
        n = total_length
    }

    // Initialize output slice of size n
    out_slice := make([]Comparable, n)

    for i := 0; i < n; i++ {
        // Find the smallest element among all slices by inspecting the first
        // element in each slice since all slices are ordered in ascending
        // order
        low := -1
        for j := 0; j < k; j++ {
            com, ok := Peek(slices[j])
            if (ok) {
                if (low == -1) {
                    low = j
                } else if (slices[low][0].CompareTo(com) > -1) {
                    low = j
                }
            }
        }
        // Append lowest element to out_slice
        out_slice[i] = slices[low][0]

        // Now 'extract' lowest element from slice by manipulating slice
        // pointers
        slices[low] = slices[low][1:]
    }

    return out_slice
}

func MergeKSortedSlicesTopNHeap(n int, slices ...[]Comparable) []Comparable {

    k := len(slices) // no. slices

    // Check for sane input 
    // if n is larger than the combined length of the slices then
    // set n equals the combined length of the slices
    total_length := 0
    for i := 0; i < k; i++ {
        total_length = total_length + len(slices[i])
    }

    if total_length < n {
        n = total_length
    }

    heapz := make(Heap, k) // new heap of size k

    for i := 0; i < k; i++ {
        a, ok := Peek(slices[i])
        if (ok) {
            heapz[i] = &HeapElement{
                       k: i,
                       elem: a,
            }
			if (len(slices[i]) > 0) {
				slices[i] = slices[i][1:]
			}
        }
    }

	heap.Init(&heapz)

    // Initialize output slice of size n
    out_slice := make([]Comparable, n)

    for i := 0; i < n; i++ {
        elem := heap.Pop(&heapz).(*HeapElement)
        k    := elem.k
        item := elem.elem

        // insert into outslice
        out_slice[i] = item
        // Extract from slice k
		a, ok := Peek(slices[k])
		if (ok) {
			elem := &HeapElement{k: k, elem: a}
			heap.Push(&heapz, elem)
			slices[k] = slices[k][1:]
		}

    }
    return out_slice
}

// Peek into a slice revealing the element at index 0 if it exists
func Peek(slice []Comparable) (Comparable, bool) {
    // Perhaps use type interface{}
    if (len(slice) > 0) {
        return slice[0], true
    }
    return nil, false
}

// Heap implementations ---------------------------------------------------
type HeapElement struct {
    k    int // Perhaps make this uint8 or uint16
    elem Comparable // Perhaps use pointers
}

type Heap []*HeapElement // TODO pointer receivers

// Interface implementation used in heap interface
func (h Heap) Len() int { return len(h) } // Use pointer-receivers
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool {
	i1 := h[i]
	i2 := h[j]
	com := i1.elem.CompareTo(i2.elem)
    if (com < 0) {
        return true
    }
    return false
}

func (h *Heap) Push(x interface{}) {
	item := x.(*HeapElement)
	*h = append(*h, item)
}

func (h *Heap) Pop() interface{} {
	n := len(*h)
	old := *h
	elem := old[n-1]
	*h = old[0:n-1]
	return elem
}
