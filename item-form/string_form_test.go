package form

import (
	"testing"
)

func TestStringForm(t *testing.T) {
	form := NewStringForm("name")

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"空文字が許可されること", "", false},
		{"通常の文字列が許可されること", "John Doe", false},
		{"特殊文字が許可されること", "!@#$%", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := form.Validate(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("StringForm.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
