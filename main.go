package main

import (
	"fmt"
	form "go_generics/item-form"
)

func main() {
	// フォームの作成
	forms := []form.ItemForm{
		// 文字列フォーム
		form.NewStringForm("username").
			SetValidation([]string{"required", "max:50"}).
			SetValue("John Doe"),

		// 数値フォーム
		form.NewNumberForm("age").
			SetValidation([]string{"required", "min:0", "max:150"}).
			SetValue("25"),

		// 選択フォーム
		form.NewSingleForm("role", []string{"admin", "user", "guest"}).
			SetValidation([]string{"required"}).
			SetValue("user"),
	}

	// フォームの型と値を表示
	fmt.Println("=== Form Values ===")
	for _, f := range forms {
		switch v := f.(type) {
		case *form.StringForm:
			fmt.Printf("String[%s]: %s\n", v.ID(), v.Value())
		case *form.NumberForm:
			fmt.Printf("Number[%s]: %s\n", v.ID(), v.Value())
		case *form.SingleForm:
			fmt.Printf("Single[%s]: %s (choices: %v)\n", v.ID(), v.Value(), v.Candidates())
		}
	}

	// バリデーションのテスト
	fmt.Println("\n=== Validation Test ===")
	testValues := map[string]string{
		"username": "Jane Smith",
		"age":      "invalid",
		"role":     "superuser",
	}

	for _, f := range forms {
		value := testValues[f.ID()]
		if err := f.Validate(value); err != nil {
			fmt.Printf("❌ %s: %v\n", f.ID(), err)
		} else {
			fmt.Printf("✅ %s: %s\n", f.ID(), f.Value())
		}
	}
}
