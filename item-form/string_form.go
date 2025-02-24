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
