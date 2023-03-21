package volume

import (
	"math"
)

func Cylinder(length, radius, depth float64) float64 {
	part1 := (math.Pow(radius, 2)) * (math.Acos((radius - depth) / radius))
	part2 := (radius - depth) * (math.Sqrt((2 * radius * depth) - math.Pow(depth, 2)))

	return length * (part1 - part2)
}
