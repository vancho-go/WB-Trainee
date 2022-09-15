// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func newPoint(x1 int, y1 int) *point {
	return &point{x: x1, y: y1}
}

func findDistance(p1 *point, p2 *point) float64 {
	xD, yD := p1.x-p2.x, p1.y-p2.y
	xS, yS := xD*xD, yD*yD
	res := math.Sqrt(float64(xS + yS))
	return res

}

func main() {
	p1 := newPoint(5, 8)
	p2 := newPoint(9, 3)
	fmt.Println(findDistance(p1, p2))
}
