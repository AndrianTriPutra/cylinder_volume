package volume

import (
	"errors"
	"math"
)

// Cylinder menghitung volume air dalam tangki silinder horizontal.
// length: panjang tabung (meter)
// radius: jari-jari tabung (meter)
// depth: kedalaman air dari dasar tabung (meter)
// Mengembalikan volume dalam meter kubik dan error jika input tidak valid.
func Cylinder(length, radius, depth float64) (float64, error) {
	// Validasi nilai input
	if radius <= 0 {
		return 0, errors.New("radius harus lebih dari 0")
	}
	if length <= 0 {
		return 0, errors.New("length harus lebih dari 0")
	}
	if depth < 0 || depth > 2*radius {
		return 0, errors.New("depth harus berada di antara 0 dan 2 * radius")
	}

	// Kasus ekstrem
	if depth == 0 {
		return 0, nil // kosong
	}
	if depth == 2*radius {
		fullVolume := math.Pi * radius * radius * length
		return fullVolume, nil // penuh
	}

	// Amankan argumen untuk Acos
	cosArg := (radius - depth) / radius
	if cosArg < -1 {
		cosArg = -1
	}
	if cosArg > 1 {
		cosArg = 1
	}

	// Hitung luas irisan dan volume
	part1 := math.Pow(radius, 2) * math.Acos(cosArg)
	part2 := (radius - depth) * math.Sqrt((2*radius*depth)-math.Pow(depth, 2))
	segmentArea := part1 - part2
	volume := segmentArea * length

	return volume, nil
}
