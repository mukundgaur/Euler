package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	ImplicitEuler(0.1)
}

func Euler(step float64) {
	t := 0.0
	var y float64 = 2.0
	precision := 0.001
	for t+precision < 1.99 {
		fmt.Println("y: " + strconv.FormatFloat(y, 'G', 10, 64))
		fmt.Println(5*t - 3*math.Sqrt(y))
		fmt.Println(t)
		y += (5*t - 3*math.Sqrt(y)) * step
		t += step

	}
	fmt.Println("y: " + strconv.FormatFloat(y, 'G', 10, 64))
	fmt.Println(5*t - 3*math.Sqrt(y))
}

func ImplicitEuler(step float64) {
	t := 0.0
	var y float64 = 2.0
	precision := 0.001
	for t+precision < 0.4 {
		y = (y + 2*step - step*(t+step)) / (1 - step)
		fmt.Println(y)
		t += step

	}
	fmt.Println(t)
	fmt.Println("y: " + strconv.FormatFloat(y, 'G', 10, 64))
}
