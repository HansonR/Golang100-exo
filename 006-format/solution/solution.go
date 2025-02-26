package main

import "fmt"

func main() {

	var name string
	var age int64
	var legal bool
	var weight float32

	name = "Jason"
	age = 25
	legal = true
	weight = 78.12

	fmt.Printf("My name is %s, I am %d years old and it's %t that I can drive a car, my pet weights %f kilograms", name, age, legal, weight)
}

// rappel : https://pkg.go.dev/fmt#hdr-Printing
