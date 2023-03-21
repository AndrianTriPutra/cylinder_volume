package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"cylinder-volume/volume"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("      go run . length radius depth")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	length, err := strconv.ParseFloat(flag.Args()[0], 64)
	if err != nil {
		log.Fatalf("error convert length cause:%s", err.Error())
	}

	radius, err := strconv.ParseFloat(flag.Args()[1], 64)
	if err != nil {
		log.Fatalf("error convert radius cause: %s", err.Error())
	}

	depth, err := strconv.ParseFloat(flag.Args()[2], 64)
	if err != nil {
		log.Fatalf("error convert depth cause: %s", err.Error())
	}

	fmt.Printf("length: %.2f\n", length)
	fmt.Printf("radius: %.2f\n", radius)
	fmt.Printf("depth : %.2f\n", depth)
	volume := volume.Cylinder(length, radius, depth)
	fmt.Printf("volume: %.2f\n", volume)
}
