package form

import (
	"fmt"
)

type SingleForm struct {
	BaseForm
	candidates []string
}

func NewSingleForm(id string, candidates []string) *SingleForm {
	return &SingleForm{
		BaseForm: BaseForm{
			id: id,
		},
		candidates: candidates,
	}
}

func (f *SingleForm) Validate(value string) error {
	for _, candidate := range f.candidates {
		if value == candidate {
			f.value = value
			return nil
		}
	}
	return fmt.Errorf("value must be one of %v", f.candidates)
}

func (f *SingleForm) SetValue(value string) *SingleForm {
	f.value = value
	return f
}

func (f *SingleForm) SetValidation(validation []string) *SingleForm {
	f.validation = validation
	return f
}

func (f *SingleForm) Candidates() []string {
	return f.candidates
}
