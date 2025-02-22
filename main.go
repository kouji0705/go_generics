package main

import "fmt"

// 数値型制約
type Number interface {
	int | float64
}

// 文字列型制約
type StringType interface {
	string
}

// Value関数を持つ構造体
type Container[T any] struct {
	Value T
}

// 値を取得する関数
func GetValue[T any](c Container[T]) T {
	return c.Value
}

func main() {
	// 数値型
	intBox := Container[int]{Value: 100}
	fmt.Println("Int Value:", GetValue(intBox)) // 出力: Int Value: 100

	// 文字列型
	stringBox := Container[string]{Value: "Hello"}
	fmt.Println("String Value:", GetValue(stringBox)) // 出力: String Value: Hello
}
