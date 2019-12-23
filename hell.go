package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var x, y = 2, 4
var x1, y1 = float64(x), float64(y)

func print(thing string) {
	fmt.Println(thing)
}
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			tmp := strconv.Itoa(i)
			print(tmp)
		}
	}
}

/*
to do
1. learn about array slicing
2. see how pointers are implemented
3. do some struct and pseudo class work
4. do some concurrency if you want to get fancy
*/
