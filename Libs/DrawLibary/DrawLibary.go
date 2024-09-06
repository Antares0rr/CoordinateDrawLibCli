package drawlibary

import "fmt"

type Point struct {
	X, Y int
}

func draw(points []Point) { // Draws the points and the Grid
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if containsPoint(points, Point{i, j}) {
				fmt.Print("x ")
			} else {
				fmt.Print("⋅ ")
			}
		}
		fmt.Printf("\n") // New line after each row
	}
}

func containsPoint(points []Point, p Point) bool {
	for _, point := range points {
		if point.X == p.X && point.Y == p.Y {
			return true
		}
	}
	return false
}
