package main

import "fmt"

// Sum は数値スライスの合計を計算
func Sum[T int | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(ints)) // 15

	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println(Sum(floats)) // 6.6
}
