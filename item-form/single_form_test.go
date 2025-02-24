package form

import (
	"testing"
)

func TestSingleForm(t *testing.T) {
	candidates := []string{"Apple", "Banana", "Cherry"}
	form := NewSingleForm("fruit", candidates)

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"選択肢の中から値を選択できること", "Apple", false},
		{"選択肢以外の値はエラーになること", "Orange", true},
		{"空文字はエラーになること", "", true},
		{"大文字小文字は区別されること", "apple", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := form.Validate(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("SingleForm.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
