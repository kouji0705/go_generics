package main

import "fmt"

// Values はマップから全ての値を取得
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func main() {
	m := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	fmt.Println(Values(m)) // [30, 25, 35]
}
