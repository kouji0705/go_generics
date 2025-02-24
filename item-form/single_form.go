package itemform

import (
	"fmt"
	"go_generics/items"
)

type SingleForm struct {
	BaseForm
	Candidates []string
}

func NewSingleForm(candidates []string) *SingleForm {
	return &SingleForm{
		Candidates: candidates,
	}
}

func (f *SingleForm) Build() items.ItemResult {
	item := items.NewSingleItem(f.Candidates)
	if f.Value != "" {
		item.SaveDraft(f.Value)
	}
	return item
}

func (f *SingleForm) Validate(value string) error {
	for _, candidate := range f.Candidates {
		if value == candidate {
			f.Value = value
			return nil
		}
	}
	return fmt.Errorf("value must be one of %v", f.Candidates)
} 