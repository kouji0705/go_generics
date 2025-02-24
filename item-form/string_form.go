package itemform

import "go_generics/items"

type StringForm struct {
	BaseForm
}

func NewStringForm() *StringForm {
	return &StringForm{}
}

func (f *StringForm) Build() items.ItemResult {
	item := items.NewStringItem()
	if f.Value != "" {
		item.SaveDraft(f.Value)
	}
	return item
}

func (f *StringForm) Validate(value string) error {
	f.Value = value
	return nil // 文字列は特に制約なし
}
