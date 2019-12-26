package main

import (
	"fmt"
)

func Sqrt(y float64) float64 {
	z := 1.0;
	change := 1.0;
	for change > .00001 {
		z -= (z*z - y) / (2*z)
		change = (z*z - y) / (2*z)
		}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
