package items

import "fmt"

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
