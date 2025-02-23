package types

// ItemResult インターフェース
type ItemResult interface {
	FormatCheck() error
	SetValue(value string) error
	GetValue() string
}

// Container ジェネリック構造体
type Container[T ItemResult] struct {
	Item T
}

func (c *Container[T]) SetValue(value string) error {
	return c.Item.SetValue(value)
}

func (c *Container[T]) GetValue() string {
	return c.Item.GetValue()
}

func (c *Container[T]) FormatCheck() error {
	return c.Item.FormatCheck()
}
