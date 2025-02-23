package main

import (
	"fmt"
	"your-module-name/items"
	"your-module-name/types"
)

func main() {
	// ItemResultsの作成
	var itemResults types.ItemResults

	// 各種アイテムの追加
	itemResults.Add(items.NewStringItem())
	itemResults.Add(items.NewNumberItem())
	itemResults.Add(items.NewSingleItem([]string{"Apple", "Banana", "Cherry"}))

	// 全アイテムに対して値を設定
	fmt.Println("--- Setting 'Hello' to all items ---")
	if errs := itemResults.SetValueAll("Hello"); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("\n--- Current values ---")
	for i, value := range itemResults.GetValues() {
		fmt.Printf("Item[%d]: %s\n", i, value)
	}

	fmt.Println("\n--- Setting valid values ---")
	itemResults[0].SetValue("Hello")  // StringItem
	itemResults[1].SetValue("123.45") // NumberItem
	itemResults[2].SetValue("Apple")  // SingleItem

	fmt.Println("\n--- Final values ---")
	for i, value := range itemResults.GetValues() {
		fmt.Printf("Item[%d]: %s\n", i, value)
	}
}
