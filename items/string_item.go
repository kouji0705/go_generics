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

func (s *StringItem) FormatCheck() error {
	return nil // 特に制約なし
}

func (s *StringItem) SetValue(value string) error {
	s.Value = value
	return s.FormatCheck()
}

func (s *StringItem) GetValue() string {
	return s.Value
}

// RequiredCheck は必須チェックを実施します
func (s *StringItem) RequiredCheck() error {
	if s.Value == "" {
		return errors.New("value is required")
	}
	return nil
}

// Save は全てのバリデーションを実施します
func (s *StringItem) Save() error {
	if err := s.FormatCheck(); err != nil {
		return err
	}
	return s.RequiredCheck()
}

// SaveDraft はフォーマットチェックのみを実施します
func (s *StringItem) SaveDraft() error {
	return s.FormatCheck()
}
