package items

import (
	"errors"
)

// StringItem 文字列アイテムの実装
type StringItem struct {
	Value string
}

func NewStringItem() *StringItem {
	return &StringItem{}
}

// formatCheck は文字列のフォーマットチェックを実施します
func (s *StringItem) formatCheck() error {
	return nil // 特に制約なし
}

// requiredCheck は必須チェックを実施します
func (s *StringItem) requiredCheck() error {
	if s.Value == "" {
		return errors.New("value is required")
	}
	return nil
}

func (s *StringItem) SetValue(value string) error {
	s.Value = value
	return s.formatCheck()
}

func (s *StringItem) GetValue() string {
	return s.Value
}

// Save は全てのバリデーションを実施します
func (s *StringItem) Save() error {
	if err := s.formatCheck(); err != nil {
		return err
	}
	return s.requiredCheck()
}

// SaveDraft はフォーマットチェックのみを実施します
func (s *StringItem) SaveDraft() error {
	return s.formatCheck()
}
