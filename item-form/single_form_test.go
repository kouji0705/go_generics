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
		{"valid choice", "Apple", false},
		{"invalid choice", "Orange", true},
		{"empty value", "", true},
		{"case sensitive", "apple", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := form.Validate(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("SingleForm.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
} 