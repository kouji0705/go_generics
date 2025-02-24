package form

type ItemForm interface {
	Validate(value string) error
	ID() string
	Value() string
	Validation() []string
	// GetID() string
	// GetType() string
	// GetValidation() []string
	// GetCandidates() []string
}

// BaseForm は共通のフォーム機能を提供する基底構造体
type BaseForm struct {
	id         string   // フォームの一意識別子
	value      string   // フォームの値
	validation []string // バリデーションルール
	// Candidates []string // 選択肢（該当する場合のみ使用）
}

func (f *BaseForm) ID() string {
	return f.id
}

func (f *BaseForm) Value() string {
	return f.value
}

func (f *BaseForm) Validation() []string {
	return f.validation
}

// func (f *BaseForm) GetID() string {
// 	return f.ID
// }

// func (f *BaseForm) GetType() string {
// 	return f.Type
// }

// func (f *BaseForm) GetValidation() []string {
// 	return f.Validation
// }

// func (f *BaseForm) GetCandidates() []string {
// 	return f.Candidates
// }

// func (f *BaseForm) GetValue() string {
// 	return f.Value
// }
