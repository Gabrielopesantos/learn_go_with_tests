package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assertResult := func(t *testing.T, got, hasArea float64) {
		if got != hasArea {
			// t.Errorf("got %v, hasArea %v", got, hasArea)
			t.Errorf("got %.2f, hasArea %.2f", got, hasArea)
		}
	}
	t.Run("Testing perimeter", func(t *testing.T) {
		t.Helper()
		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		hasArea := 40.0
		assertResult(t, got, hasArea)
	})
}

// func TestArea(t *testing.T) {
// 	assertResult := func(t *testing.T, shape Shape, hasArea float64) {
// 		got := shape.Area()
// 		if got != hasArea {
// 			// t.Errorf("got %v, hasArea %v", got, hasArea)
// 			t.Errorf("got %g, hasArea %g", got, hasArea)
// 		}
// 	}
// 	t.Run("Testing rectangle area", func(t *testing.T) {
// 		t.Helper()
// 		rect := &Rectangle{10.0, 10.0}
// 		assertResult(t, rect, 100.0)
// 	})

// 	t.Run("Testing circle area", func(t *testing.T) {
// 		t.Helper()
// 		circle := &Circle{10.0}
// 		assertResult(t, circle, math.Pi*math.Pow(10.0, 2))
// 	})
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: &Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{10, 10}, hasArea: 50.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, hasArea %g", tt, got, tt.hasArea)
			}
		})
	}
}
