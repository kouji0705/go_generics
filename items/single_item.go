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

// formatCheck は選択肢のフォーマットチェックを実施します
func (s *SingleItem) formatCheck() error {
	for _, candidate := range s.Candidates {
		if s.Value == candidate {
			return nil
		}
	}
	return fmt.Errorf("value must be one of %v", s.Candidates)
}

// requiredCheck は必須チェックを実施します
func (s *SingleItem) requiredCheck() error {
	if s.Value == "" {
		return errors.New("value is required")
	}
	return nil
}

func (s *SingleItem) GetValue() string {
	return s.Value
}

// Save は全てのバリデーションを実施します
func (s *SingleItem) Save(value string) error {
	s.Value = value
	if err := s.formatCheck(); err != nil {
		return err
	}
	return s.requiredCheck()
}

// SaveDraft はフォーマットチェックのみを実施します
func (s *SingleItem) SaveDraft(value string) error {
	s.Value = value
	return s.formatCheck()
}
