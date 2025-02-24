package form

type StringForm struct {
	BaseForm
}

func NewStringForm(id string) *StringForm {
	return &StringForm{
		BaseForm: BaseForm{
			id: id,
		},
	}
}

func (f *StringForm) Validate(value string) error {
	f.value = value
	return nil // 文字列は特に制約なし
}

func (f *StringForm) SetValue(value string) *StringForm {
	f.value = value
	return f
}

func (f *StringForm) SetValidation(validation []string) *StringForm {
	f.validation = validation
	return f
}
