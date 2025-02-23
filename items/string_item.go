package items

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
