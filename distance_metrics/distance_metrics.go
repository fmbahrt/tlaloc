// Package provides functionality to calculate
// distances between points in a n-dimensional
// space.
// As of right now this package only support float32
package distance_metrics

import (
    "errors"
    "math"
)

// Calculates the euclidean distance between 
// two points in a n-dimensional euclidean space.
// This implementation does not includ square root.
func EuclideanDistanceNoSqrt(v1 []float32, v2 []float32) (float32, error) {
    if (len(v1) != len(v2)){
        return -1.0, errors.New("Slices are not of the same size.")
    }
    if (len(v1) == 0){
        return -1.0, errors.New("Slices cannot be of length 0.")
    }

    sum := float32(0.0) // Type inference?

    for i := 0; i < len(v1); i++{
        min := v2[i] - v1[i]
        sum += min * min
    }

    return sum, nil
}

// Calculates the manhattan distance between 
// two point(vecors) in a n-dimensional space.
func ManhattanDistance(v1 []float32, v2 []float32) (float32, error) {
    if (len(v1) != len(v2)){
        return -1.0, errors.New("Slices are not of the same size.")
    }
    if (len(v1) == 0){
        return -1.0, errors.New("Slices cannot be of length 0.")
    }

    sum := float32(0.0) //Type inference?

    for i :=0; i < len(v1); i++{
        //abs
        sum += math.Float32frombits(math.Float32bits((v1[i]-v2[i])) &^ (1 << 31))
    }

    return sum, nil
}
