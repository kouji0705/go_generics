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

// formatCheck は数値のフォーマットチェックを実施します
func (n *NumberItem) formatCheck() error {
	if n.Value < 0 {
		return errors.New("value must be a non-negative number")
	}
	return nil
}

// requiredCheck は必須チェックを実施します
func (n *NumberItem) requiredCheck() error {
	if n.Value == 0 {
		return errors.New("value is required")
	}
	return nil
}

func (n *NumberItem) SetValue(value string) error {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("invalid number format")
	}
	n.Value = num
	return n.formatCheck()
}

func (n *NumberItem) GetValue() string {
	return fmt.Sprintf("%f", n.Value)
}

// Save は全てのバリデーションを実施します
func (n *NumberItem) Save(value string) error {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("invalid number format")
	}
	n.Value = num
	if err := n.formatCheck(); err != nil {
		return err
	}
	return n.requiredCheck()
}

// SaveDraft はフォーマットチェックのみを実施します
func (n *NumberItem) SaveDraft(value string) error {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("invalid number format")
	}
	n.Value = num
	return n.formatCheck()
}
