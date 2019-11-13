package clockface

import (
	"time"
	"math"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

func SecondHand(t time.Time) Point {
	p := secondsHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}	// scale
	p = Point{p.X, -p.Y}	// flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}	// translate
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondsHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
