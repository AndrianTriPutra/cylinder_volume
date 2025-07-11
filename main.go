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
		fmt.Println("      go run . length diameter depth")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	length, err := strconv.ParseFloat(flag.Args()[0], 64)
	if err != nil {
		log.Fatalf("error convert length:%v", err)
	}

	diameter, err := strconv.ParseFloat(flag.Args()[1], 64)
	if err != nil {
		log.Fatalf("error convert diameter:%v", err)
	}

	depth, err := strconv.ParseFloat(flag.Args()[2], 64)
	if err != nil {
		log.Fatalf("error convert depth:%v", err)
	}

	radius := diameter / 2
	log.Printf("length  : %.4f", length)
	log.Printf("diameter: %.4f", diameter)
	log.Printf("radius  : %.4f", radius)
	log.Printf("depth   : %.4f", depth)

	volume, err := volume.Cylinder(length, radius, depth)
	if err != nil {
		log.Fatalf("Failed Calculate Volume:%v", err)
	}
	log.Printf("Volume  : %.4f (m3)", volume)

}
