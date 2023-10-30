package main

import "testing"

// func TestPerimeter(t *testing.T) {

// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T) {

// 	got := Area(10.0, 10.0)
// 	want := 100.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// Refactor add struct
func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Area(rectangle)
// 	want := 100.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}

// }

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {

// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Area()
// 		want := 100.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}

// 	})

// 	t.Run("circles", func(t *testing.T) {

// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

// Refactor

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shap Shape, want float64) {
// 		t.Helper()
// 		got := shap.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// func testArea(t *testing.T) {
// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Rectangle{12, 6}, 72.0},
// 		{Circle{10}, 314.1592653589793},
// 		// add a triangle
// 		{Triangle{12, 6}, 36.0},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %g want %g", got, tt.want)
// 		}
// 	}
// }

func testArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		// add a triangle
		{name: "Triangle", shape: Triangle{Base: 12, Height: 12}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
