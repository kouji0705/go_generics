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
		{"valid integer", "123", false},
		{"valid decimal", "123.45", false},
		{"invalid number", "abc", true},
		{"empty value", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := form.Validate(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("NumberForm.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
