package itemform

import "go_generics/items"

// Form はフォームの基底インターフェース
type Form interface {
	Build() items.ItemResult
	Validate(value string) error
}

// BaseForm は共通のフォーム機能を提供する基底構造体
type BaseForm struct {
	Value string
}

func (f *BaseForm) GetValue() string {
	return f.Value
} 