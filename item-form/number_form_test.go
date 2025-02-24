package form

import (
	"testing"
)

func TestNumberForm(t *testing.T) {
	form := NewNumberForm("age")

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"整数値が正しく検証できること", "123", false},
		{"小数値が正しく検証できること", "123.45", false},
		{"数値以外の文字列はエラーになること", "abc", true},
		{"空文字はエラーになること", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := form.Validate(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("NumberForm.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
