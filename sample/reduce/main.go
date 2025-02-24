package main

import "fmt"

// Reduce はスライスを1つの値に集約
func Reduce[T any, U any](slice []T, initial U, accumulator func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(numbers, 0, func(acc int, n int) int { return acc + n })
	fmt.Println(sum) // 15

	words := []string{"Go", "is", "awesome"}
	concat := Reduce(words, "", func(acc string, s string) string { return acc + s + " " })
	fmt.Println(concat) // "Go is awesome "
}
