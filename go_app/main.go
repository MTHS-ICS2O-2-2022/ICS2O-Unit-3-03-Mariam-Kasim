// Created by: Mariam Kasim
// Created on: Apr 2023
//
// This program finds the volume of a sphere

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function finds the volume of a sphere
	var radius float64

	// input
	fmt.Println("This program finds the volume of a sphere.")
	fmt.Println()
	fmt.Print("Enter the radius (in): ")
	fmt.Scanln(&radius)

	// process
	volume := (4.0 / 3.0) * math.Pi * (math.Pow(radius, 3))
	roundVolume := fmt.Sprintf("%.2f", volume)

	// output
	fmt.Println()
	fmt.Println("The volume is: " + roundVolume + " in³")
	fmt.Println("\nDone.")
}
