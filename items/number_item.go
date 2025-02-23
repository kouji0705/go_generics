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
