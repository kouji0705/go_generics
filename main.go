package main

import (
	"fmt"
	"your-module-name/items"
	"your-module-name/types"
)

func main() {
	// StringItemの例
	stringContainer := types.Container[types.ItemResult]{
		Item: items.NewStringItem(),
	}
	if err := stringContainer.SetValue("Hello"); err != nil {
		fmt.Println("StringItem Error:", err)
	} else {
		fmt.Println("StringItem Value:", stringContainer.GetValue())
	}

	// NumberItemの例
	numberContainer := types.Container[types.ItemResult]{
		Item: items.NewNumberItem(),
	}
	if err := numberContainer.SetValue("123.45"); err != nil {
		fmt.Println("NumberItem Error:", err)
	} else {
		fmt.Println("NumberItem Value:", numberContainer.GetValue())
	}

	// SingleItemの例
	singleContainer := types.Container[types.ItemResult]{
		Item: items.NewSingleItem([]string{"Apple", "Banana", "Cherry"}),
	}
	if err := singleContainer.SetValue("Banana"); err != nil {
		fmt.Println("SingleItem Error:", err)
	} else {
		fmt.Println("SingleItem Value:", singleContainer.GetValue())
	}

	// バリデーション失敗例
	if err := singleContainer.SetValue("Orange"); err != nil {
		fmt.Println("Invalid SingleItem:", err)
	}
}
