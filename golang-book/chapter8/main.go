package main

import (
	"fmt"
	m "github.com/djgriffin/learning/golang-book/chapter8/math"
)

func main() {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
