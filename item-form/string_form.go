package form

type StringForm struct {
	BaseForm
}

func NewStringForm(id string, validation []string) *StringForm {
	return &StringForm{
		BaseForm: BaseForm{
			ID:   id,
			Type: "string",
			// Validation: validation,
		},
	}
}

func (f *StringForm) Validate(value string) error {
	f.Value = value
	return nil // 文字列は特に制約なし
}
