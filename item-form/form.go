package form

// Form はフォームの基底インターフェース
type Form interface {
	// Build() items.ItemResult
	Validate(value string) error
	GetID() string
	GetType() string
	GetValidation() []string
	GetCandidates() []string
}

// BaseForm は共通のフォーム機能を提供する基底構造体
type BaseForm struct {
	ID         string   // フォームの一意識別子
	Type       string   // フォームの種類（string, number, single など）
	Value      string   // フォームの値
	Validation []string // バリデーションルール
	Candidates []string // 選択肢（該当する場合のみ使用）
}

func (f *BaseForm) GetID() string {
	return f.ID
}

func (f *BaseForm) GetType() string {
	return f.Type
}

func (f *BaseForm) GetValidation() []string {
	return f.Validation
}

func (f *BaseForm) GetCandidates() []string {
	return f.Candidates
}

func (f *BaseForm) GetValue() string {
	return f.Value
}
