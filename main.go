package main

import (
	"fmt"
	"go_generics/items"
)

func main() {
	// ItemResultsの作成
	var itemResults items.ItemResults

	// 各種アイテムの追加
	itemResults.Add(items.NewStringItem())
	itemResults.Add(items.NewNumberItem())
	itemResults.Add(items.NewSingleItem([]string{"Apple", "Banana", "Cherry"}))

	// 全アイテムに対して値を設定
	fmt.Println("--- Setting 'Hello' to all items ---")
	if errs := itemResults.SaveDraftAll("Hello"); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("\n--- Current values ---")
	for i, value := range itemResults.GetValues() {
		fmt.Printf("Item[%d]: %s\n", i, value)
	}

	fmt.Println("\n--- Setting valid values ---")
	itemResults[0].SaveDraft("Hello")  // StringItem
	itemResults[1].SaveDraft("123.45") // NumberItem
	itemResults[2].SaveDraft("Apple")  // SingleItem

	fmt.Println("\n--- Final values ---")
	for i, value := range itemResults.GetValues() {
		fmt.Printf("Item[%d]: %s\n", i, value)
	}

	// SaveDraftのテスト
	fmt.Println("\n--- Testing SaveDraft ---")
	for i, item := range itemResults {
		if err := item.SaveDraft("test"); err != nil {
			fmt.Printf("Item[%d] SaveDraft Error: %v\n", i, err)
		} else {
			fmt.Printf("Item[%d] SaveDraft OK\n", i)
		}
	}

	// Saveのテスト
	fmt.Println("\n--- Testing Save ---")
	for i, item := range itemResults {
		if err := item.Save("test"); err != nil {
			fmt.Printf("Item[%d] Save Error: %v\n", i, err)
		} else {
			fmt.Printf("Item[%d] Save OK\n", i)
		}
	}
}
