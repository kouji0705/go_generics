package types

import (
	"fmt"
)

// ItemResult インターフェース
type ItemResult interface {
	FormatCheck() error
	SetValue(value string) error
	GetValue() string
}

// Container ジェネリック構造体
type Container[T ItemResult] struct {
	Item T
}

func (c *Container[T]) SetValue(value string) error {
	return c.Item.SetValue(value)
}

func (c *Container[T]) GetValue() string {
	return c.Item.GetValue()
}

func (c *Container[T]) FormatCheck() error {
	return c.Item.FormatCheck()
}

// ItemResults は複数のItemResultを管理するためのスライス型
type ItemResults []ItemResult

// Add は新しいItemResultを追加します
func (items *ItemResults) Add(item ItemResult) {
	*items = append(*items, item)
}

// SetValueAll は全てのアイテムに対して値を設定します
func (items ItemResults) SetValueAll(value string) []error {
	var errors []error
	for i, item := range items {
		if err := item.SetValue(value); err != nil {
			errors = append(errors, fmt.Errorf("item[%d]: %w", i, err))
		}
	}
	return errors
}

// GetValues は全てのアイテムの値を取得します
func (items ItemResults) GetValues() []string {
	values := make([]string, len(items))
	for i, item := range items {
		values[i] = item.GetValue()
	}
	return values
}
