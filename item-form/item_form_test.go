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

// func TestFormGetters(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		form       ItemForm
// 		wantID     string
// 		wantType   string
// 		candidates []string
// 	}{
// 		{
// 			name:     "string form",
// 			form:     NewStringForm("name"),
// 			wantID:   "name",
// 			wantType: "string",
// 		},
// 		{
// 			name:     "number form",
// 			form:     NewNumberForm("age"),
// 			wantID:   "age",
// 			wantType: "number",
// 		},
// 		{
// 			name:       "single form",
// 			form:       NewSingleForm("fruit", []string{"Apple", "Banana"}),
// 			wantID:     "fruit",
// 			wantType:   "single",
// 			candidates: []string{"Apple", "Banana"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.form.GetID(); got != tt.wantID {
// 				t.Errorf("GetID() = %v, want %v", got, tt.wantID)
// 			}
// 			if got := tt.form.GetType(); got != tt.wantType {
// 				t.Errorf("GetType() = %v, want %v", got, tt.wantType)
// 			}
// 			if tt.candidates != nil {
// 				if got := tt.form.GetCandidates(); len(got) != len(tt.candidates) {
// 					t.Errorf("GetCandidates() = %v, want %v", got, tt.candidates)
// 				}
// 			}
// 		})
// 	}
// }
