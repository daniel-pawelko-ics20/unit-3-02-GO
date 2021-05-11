// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that volume of a pyramid

package main

import (
	"fmt"
	"math"
)

// This main function will ask user for dimensions and output calculations
func main() {
	// Defining variables
	var length float64
	var width float64
	var height float64

	fmt.Println("Volume of ")
	fmt.Println("A = (length x width x height) / 3")
	fmt.Println()

	// User Input
	fmt.Println("Please enter dimensions:")
	fmt.Println("Length(mm): ")
	fmt.Scanln(&length)
	fmt.Println()

	fmt.Println("Width(mm):")
	fmt.Scanln(&width)
	fmt.Println()

	fmt.Println("height(mm):")
	fmt.Scanln(&height)
	fmt.Println()

	// calculations
	var volume float64 = (length * width * height) / 3

	// Print out volume
	fmt.Println("Volume is: ", math.Round(volume*100)/100, "mm³")
}
