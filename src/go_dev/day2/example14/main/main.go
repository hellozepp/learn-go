package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
		fmt.Println(math.Pi)
	}

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
	fmt.Println(split(17))
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
