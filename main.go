package main

import (
	"errors"
	"fmt"
)

// 数値型制約
type Number interface {
	int | float64
}

// 文字列型制約
type StringType interface {
	string
}

// Container構造体
type Container[T any] struct {
	Value T
}

// バリデーション関数
func Validate[T any](c Container[T]) error {
	switch v := any(c.Value).(type) {
	case int:
		if v <= 0 {
			return errors.New("int value must be greater than 0")
		}
	case float64:
		if v <= 0.0 {
			return errors.New("float64 value must be greater than 0")
		}
	case string:
		if v == "" {
			return errors.New("string value cannot be empty")
		}
	default:
		return errors.New("unsupported type")
	}
	return nil
}

// 値を取得する関数
func GetValue[T any](c Container[T]) T {
	return c.Value
}

func main() {
	// 有効な例
	intBox := Container[int]{Value: 100}
	if err := Validate(intBox); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Int Value:", GetValue(intBox)) // 出力: Int Value: 100
	}

	stringBox := Container[string]{Value: "Hello"}
	if err := Validate(stringBox); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("String Value:", GetValue(stringBox)) // 出力: String Value: Hello
	}

	// 無効な例
	invalidInt := Container[int]{Value: -1}
	if err := Validate(invalidInt); err != nil {
		fmt.Println("Validation failed:", err) // 出力: Validation failed: int value must be greater than 0
	}

	invalidString := Container[string]{Value: ""}
	if err := Validate(invalidString); err != nil {
		fmt.Println("Validation failed:", err) // 出力: Validation failed: string value cannot be empty
	}
}
