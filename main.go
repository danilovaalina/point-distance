package main

import (
	"fmt"
	"math"
)

// Point представляет точку на плоскости.
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки.
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance вычисляет евклидово расстояние между текущей точкой и другой точкой other.
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	dist := p1.Distance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", dist) // Выведет: 5.00
}
