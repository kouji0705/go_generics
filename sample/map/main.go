package main

import "fmt"

// Keys はマップから全てのキーを取得
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	m := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	fmt.Println(Keys(m)) // ["Alice", "Bob", "Charlie"]
}
