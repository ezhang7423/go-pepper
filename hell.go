package main

import (
	"fmt"
	"math/rand"
)

var x, y = 2, 4
var x1, y1 = float64(x), float64(y)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
