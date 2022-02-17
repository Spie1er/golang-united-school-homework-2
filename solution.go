package square

import (
	"math"
)

type Side int

const (
	SidesCircle   Side = 0
	SidesTriangle Side = 3
	SidesSquare  Side = 4
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)


func CalcSquare(sideLen float64, sidesNum Side) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4 * sideLen * sideLen
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else {
		return 0
	}
}