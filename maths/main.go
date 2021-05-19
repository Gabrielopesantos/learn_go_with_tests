package clockface

import (
    "time"
    "math"
)

type Point struct {
    X float64
    Y float64
}

func SecondHand(t time.Time) Point {
    return Point{}
}

func SecondsInRadians(t time.Time) float64 {
    // return float64(t.Second()) * (math.Pi * 30) 
    return math.Pi / (30 / float64(t.Second()))
}

func SecondHandPoint(t time.Time) Point {
    angle := SecondsInRadians(t)
    x := math.Sin(angle)
    y := math.Cos(angle)

    return Point{x, y}
}

func roughlyEqualFloat64(a, b float64) bool {
    const equalityThreshold = 1e-7
    return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
    return roughlyEqualFloat64(a.X, b.X) && 
        roughlyEqualFloat64(a.Y, b.Y)
    }
