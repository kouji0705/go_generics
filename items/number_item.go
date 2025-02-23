package items

import (
	"errors"
	"fmt"
	"strconv"
)

type NumberItem struct {
	Value float64
}

func NewNumberItem() *NumberItem {
	return &NumberItem{}
}

func (n *NumberItem) FormatCheck() error {
	if n.Value < 0 {
		return errors.New("value must be a non-negative number")
	}
	return nil
}

func (n *NumberItem) SetValue(value string) error {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("invalid number format")
	}
	n.Value = num
	return n.FormatCheck()
}

func (n *NumberItem) GetValue() string {
	return fmt.Sprintf("%f", n.Value)
}

// RequiredCheck は必須チェックを実施します
func (n *NumberItem) RequiredCheck() error {
	if n.Value == 0 {
		return errors.New("value is required")
	}
	return nil
}

// Save は全てのバリデーションを実施します
func (n *NumberItem) Save() error {
	if err := n.FormatCheck(); err != nil {
		return err
	}
	return n.RequiredCheck()
}

// SaveDraft はフォーマットチェックのみを実施します
func (n *NumberItem) SaveDraft() error {
	return n.FormatCheck()
}
