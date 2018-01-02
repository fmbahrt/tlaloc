package distance_metrics //perhaps make own package for tests

import "testing"

//t.Errorf
//t.Fatalf - stops execution

// Use testing table (tt)

// Use Subtest

var euclidDistNoSqrtTest = []struct{
    name string
    vector1 []float32
    vector2 []float32
    distance float32
    errors bool
}{
    {"Identical vectors",
     []float32{1.5, 2.0, 3.5},
     []float32{1.5, 2.0, 3.5},
     float32(0.0), false},
    {"Vectors of size 1",
     []float32{1.0},
     []float32{2.0},
     float32(1.0), false},
    {"Empty vectors",
     []float32{},
     []float32{},
     float32(-1.0), true},
    {"Vectors not of the same size",
     []float32{1.0, 2.0},
     []float32{1.0},
     float32(-1.0), true},
    {"Vectors not identical",
     []float32{1.0, 2.0, 3.0, 4.0, 5.0},
     []float32{5.0, 4.0, 3.0, 8.0, 16.0},
     float32(157), false},
}

func TestEuclideanDistance(t *testing.T){
    for _, tt := range euclidDistNoSqrtTest{
        t.Run(tt.name, func(t *testing.T){
            dist, err := EuclideanDistanceNoSqrt(tt.vector1,
                                                 tt.vector2)

            if (tt.errors == true){
                if (err == nil){ // TODO This can be done better
                    t.Fatal("Error supposed to be thrown.")
                }
            } else if (tt.distance != dist){
                t.Fatalf("Distance: %f is not %f", dist, tt.distance)
            }
        })
    }
}

var manhattanDistTest = []struct{
    name string
    vector1 []float32
    vector2 []float32
    distance float32
    errors bool
}{
    {"Identical vectors",
     []float32{1.5, 2.0, 3.5},
     []float32{1.5, 2.0, 3.5},
     float32(0.0), false},
    {"Vectors of size 1",
     []float32{1.0},
     []float32{2.0},
     float32(1.0), false},
    {"Empty vectors",
     []float32{},
     []float32{},
     float32(-1.0), true},
    {"Vectors not of the same size",
     []float32{1.0, 2.0},
     []float32{1.0},
     float32(-1.0), true},
    {"Vectors not identical",
     []float32{1.0, 2.0, 3.0, 4.0, 5.0},
     []float32{5.0, 4.0, 3.0, 8.0, 16.0},
     float32(21), false},
}

func TestManhattanDistance(t *testing.T){
    for _, tt := range manhattanDistTest{
        t.Run(tt.name, func(t *testing.T){
            dist, err := ManhattanDistance(tt.vector1,
                                           tt.vector2)

            if (tt.errors == true){
                if (err == nil){ // TODO This can be done better
                    t.Fatal("Error supposed to be thrown.")
                }
            } else if (tt.distance != dist){
                t.Fatalf("Distance: %f is not %f", dist, tt.distance)
            }
        })
    }
}
