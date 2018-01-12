package slice_utils

import (
    "testing"
)

// Test data structures for testing
type TestEntry struct {
    ID   int
    Dist float32
}

// Implement comparable interface
func (te TestEntry) CompareTo(to Comparable) int {
    x, _ := to.(TestEntry)
    if (te.Dist < x.Dist) {
        return -1
    } else if (te.Dist > x.Dist) {
        return 1
    } else {
        return 0
    }
}

var mergeKSortedListsTest = []struct{
    name string
    n    int
    slice1 []Comparable
    slice2 []Comparable
    slice3 []Comparable
    result []Comparable
}{
    {
        "3 sorted slices of equal len (3). n = 3",
        3,
        []Comparable{
            TestEntry{0, 5.4},
            TestEntry{1, 6.4},
            TestEntry{2, 7.4},
        },
        []Comparable{
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
            TestEntry{5, 9.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{7, 1.6},
            TestEntry{8, 7.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{7, 1.6},
            TestEntry{3, 2.4},
        },
    },
    {
        "3 sorted slices of equal len (3). n = 9",
        9,
        []Comparable{
            TestEntry{0, 5.4},
            TestEntry{1, 6.4},
            TestEntry{2, 7.2},
        },
        []Comparable{
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
            TestEntry{5, 9.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{7, 1.6},
            TestEntry{8, 7.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{7, 1.6},
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
            TestEntry{0, 5.4},
            TestEntry{1, 6.4},
            TestEntry{2, 7.2},
            TestEntry{8, 7.4},
            TestEntry{5, 9.4},
        },
    },
    {
        "3 sorted slices of not equal len. n = 3",
        3,
        []Comparable{
            TestEntry{0, 5.4},
            TestEntry{2, 7.4},
        },
        []Comparable{
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
            TestEntry{5, 9.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{8, 7.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
        },
    },
    {
        "3 sorted slices of not equal len. n = 9 (total)",
        9,
        []Comparable{
            TestEntry{0, 5.4},
            TestEntry{2, 7.2},
        },
        []Comparable{
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
            TestEntry{5, 9.4},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{7, 1.6},
        },
        []Comparable{
            TestEntry{6, 1.4},
            TestEntry{7, 1.6},
            TestEntry{3, 2.4},
            TestEntry{4, 3.4},
            TestEntry{0, 5.4},
            TestEntry{2, 7.2},
            TestEntry{5, 9.4},
        },
    },
//    { TODO FIX THIS CASE - very unlikely though
//       "3 sorted slices len (0). n = 3",
//        3,
//        []Comparable{},
//        []Comparable{},
//        []Comparable{},
//        []Comparable{},
//   },
}

func TestMergeKSortedSlicesTopN(t *testing.T) {
    for _, tt := range mergeKSortedListsTest {
        t.Run(tt.name, func(t *testing.T) {

            output := MergeKSortedSlicesTopN(tt.n,
                                             tt.slice1,
                                             tt.slice2,
                                             tt.slice3)

            // test if output == tt.result
            if (len(output) != len(tt.result)) {
                t.Fatalf("%v != %v", output, tt.result)
            }

            for i := 0; i < len(output); i++ {
                if ((output[i].CompareTo(tt.result[i])) != 0) {
                    t.Fatalf("%v != %v", output[i], tt.result[i])
                }
            }
        })
    }
}

func TestMergeKSortedSlicesTopNHeap(t *testing.T) {
    for _, tt := range mergeKSortedListsTest {
        t.Run(tt.name, func(t *testing.T) {

            output := MergeKSortedSlicesTopNHeap(tt.n,
                                                 tt.slice1,
                                                 tt.slice2,
                                                 tt.slice3)

            // test if output == tt.result
            if (len(output) != len(tt.result)) {
                t.Fatalf("%v != %v", output, tt.result)
            }

            for i := 0; i < len(output); i++ {
                if ((output[i].CompareTo(tt.result[i])) != 0) {
                    t.Fatalf("%v != %v", output[i], tt.result[i])
                }
            }
        })
    }
}

