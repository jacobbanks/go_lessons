package stucts_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	if got != want {
		t.Fatalf("got %.2f, wanted %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Fatalf("got %g, wanted %g", got, want)
		}
	}
	t.Run("Tests for area of a rectangle", func(t *testing.T) {
		val := Rectangle{5.0, 5.0}
		checkArea(t, val, 25.0)
	})
	t.Run("Tests for area of a circle", func(t *testing.T) {
		val := Circle{10.0}
		checkArea(t, val, 314.1592653589793)
	})
}

func TestAreaUsingTables(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{5.0, 5.0}, expected: 25.0},
		{name: "Circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.shape.Area()
			w := tt.expected
			if a != w {
				t.Fatalf("%#v got %g, wanted %g", tt.shape, a, w)
			}
		})
	}
}
