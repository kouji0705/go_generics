package items

import (
	"errors"
	"fmt"
)

type SingleItem struct {
	Value      string
	Candidates []string
}

func NewSingleItem(candidates []string) *SingleItem {
	return &SingleItem{
		Candidates: candidates,
	}
}

func (s *SingleItem) FormatCheck() error {
	for _, candidate := range s.Candidates {
		if s.Value == candidate {
			return nil
		}
	}
	return fmt.Errorf("value must be one of %v", s.Candidates)
}

func (s *SingleItem) SetValue(value string) error {
	s.Value = value
	return s.FormatCheck()
}

func (s *SingleItem) GetValue() string {
	return s.Value
}

// RequiredCheck は必須チェックを実施します
func (s *SingleItem) RequiredCheck() error {
	if s.Value == "" {
		return errors.New("value is required")
	}
	return nil
}

// Save は全てのバリデーションを実施します
func (s *SingleItem) Save() error {
	if err := s.FormatCheck(); err != nil {
		return err
	}
	return s.RequiredCheck()
}

// SaveDraft はフォーマットチェックのみを実施します
func (s *SingleItem) SaveDraft() error {
	return s.FormatCheck()
}
