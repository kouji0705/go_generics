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
		{"empty value", "", false},
		{"normal string", "John Doe", false},
		{"special chars", "!@#$%", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := form.Validate(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("StringForm.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
} 