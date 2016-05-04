package main

import (
	"fmt"
	"math"
)

type Magnituder interface {
	Magnitude() float64
}

type Vec2f struct {
	x, y float64
}

func (myVec Vec2f) Magnitude() float64 {
	return math.Sqrt(myVec.x*myVec.x + myVec.y*myVec.y)
}

// BEGIN OMIT
func DoNotDoThis(empty interface{}) {
	fmt.Println(empty.Magnitude())
}

func ThisIsAcceptable(magVec Magnituder) {
	fmt.Println(magVec.Magnitude())
}

func main() {
	myVec := Vec2f{1.0, 2.0}
	DoNotDoThis(myVec)
	ThisIsAcceptable(myVec)
}

// END OMIT
