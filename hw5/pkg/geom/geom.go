package geom

import (
	"go-search/hw5/pkg/point"
	"math"
)

func Distance(p1, p2 *point.Point) float64 {
	distance := math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
	return distance
}
