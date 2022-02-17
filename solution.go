package square

import "math"

type Side int

const (
	SideCircle   Side = 0
	SideTriangle Side = 3
	SidesSquare  Side = 4
)

func CalcSquare(sideLen float64, sidesNum Side) float64 {
	if sidesNum == SideCircle {
		return math.Pi * sideLen * sideLen
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SideTriangle {
		return math.Sqrt(3) / 4 * sideLen * sideLen
	} else {
		return 0
	}

}
