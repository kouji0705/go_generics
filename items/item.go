package items

import (
	"fmt"
)

// ItemResult インターフェース
type ItemResult interface {
	Save(value string) error
	SaveDraft(value string) error
	GetValue() string
}

// ItemResults は複数のItemResultを管理するためのスライス型
type ItemResults []ItemResult

// Add は新しいItemResultを追加します
func (items *ItemResults) Add(item ItemResult) {
	*items = append(*items, item)
}

// SaveAll は全てのアイテムに対して値を保存します
func (items ItemResults) SaveAll(value string) []error {
	var errors []error
	for i, item := range items {
		if err := item.Save(value); err != nil {
			errors = append(errors, fmt.Errorf("item[%d]: %w", i, err))
		}
	}
	return errors
}

// SaveDraftAll は全てのアイテムに対してドラフトとして保存します
func (items ItemResults) SaveDraftAll(value string) []error {
	var errors []error
	for i, item := range items {
		if err := item.SaveDraft(value); err != nil {
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
