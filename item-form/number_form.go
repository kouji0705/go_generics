package form

import (
	"strconv"
)

type NumberForm struct {
	BaseForm
}

func NewNumberForm(id string, validation []string) *NumberForm {
	return &NumberForm{
		BaseForm: BaseForm{
			ID:         id,
			Type:       "number",
			Validation: validation,
		},
	}
}

func (f *NumberForm) Validate(value string) error {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	f.Value = value
	return nil
}
