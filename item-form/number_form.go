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
