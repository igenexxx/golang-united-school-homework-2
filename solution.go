package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myInt int

const (
	triangle = 3
	square   = 4
	circle   = 0
)

func squareSquare(sideLen float64) float64 {
	return sideLen * sideLen
}

func triangleSquare(sideLen float64) float64 {
	return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
}

func circleSquare(sideLen float64) float64 {
	return math.Pi * math.Pow(sideLen, 2)
}

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	switch sidesNum {
	case square:
		squareSquare(sideLen)
	case triangle:
		triangleSquare(sideLen)
	case circle:
		circleSquare(sideLen)
	default:
		return 0
	}
}
