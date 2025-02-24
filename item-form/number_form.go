package form

import (
	"strconv"
)

type NumberForm struct {
	BaseForm
}

func NewNumberForm(id string) *NumberForm {
	return &NumberForm{
		BaseForm: BaseForm{
			id: id,
		},
	}
}

func (f *NumberForm) Validate(value string) error {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	f.value = value
	return nil
}

func (f *NumberForm) SetValue(value string) *NumberForm {
	f.value = value
	return f
}

func (f *NumberForm) SetValidation(validation []string) *NumberForm {
	f.validation = validation
	return f
}
