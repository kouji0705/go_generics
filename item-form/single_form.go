package form

import (
	"fmt"
)

type SingleForm struct {
	BaseForm
	Candidates []string
}

func NewSingleForm(id string, validation []string, candidates []string) *SingleForm {
	return &SingleForm{
		BaseForm: BaseForm{
			ID:   id,
			Type: "single",
			// Validation: validation,
			// Candidates: candidates,
		},
	}
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
