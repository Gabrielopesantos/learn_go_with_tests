package clockface

import (
    "testing"
    "time"
    "math"
 
)

func TestSecondHandAtMidnight(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

    want := Point{X: 150, Y: 150 - 90}
    got := SecondHand(tm)

    if got != want {
        t.Errorf("Got %v, wanted %v", got, want)
    }
}

func TestSecondsInRadians(t *testing.T) {
    cases := []struct{
        time time.Time
        angle float64
    }{
        {simpleTime(0, 0, 30), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(0, 0, 45), (math.Pi/2) * 3},
        {simpleTime(0, 0, 7), (math.Pi/30) * 7},
    }
    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := SecondsInRadians(c.time)

            if c.angle != got {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}

func simpleTime(hours, minutes, seconds int) time.Time {
    return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
    return t.Format("15:04:05")
}

func TestSecondHandVector(t *testing.T) {
    cases := []struct {
        time time.Time
        point Point
    }{
        {simpleTime(0, 0, 30), Point{0, -1}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := SecondHandPoint(c.time)
            if got != c.point {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}



func TestSecondHandPoint(t *testing.T) {
    cases := []struct {
        time time.Time
        point Point
    }{
        {simpleTime(0, 0, 30), Point{0, -1}},
        {simpleTime(0, 0, 45), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := SecondHandPoint(c.time)
            if got != c.point {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}
