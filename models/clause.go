package models

type IClause interface {
	GetField() string
	SetField(string)
	GetOperator() string
	SetOperator(string)
	GetValue() any
	SetValue(any)
	GetChildren() []IClause
	SetChildren([]IClause)
}

type Clause struct {
	Field    string    `json:"field"`
	Operator string    `json:"operator"`
	Value    any       `json:"value"`
	Children []IClause `json:"children"`
}

func (c *Clause) GetField() string {
	return c.Field
}

func (c *Clause) SetField(field string) {
	c.Field = field
}

func (c *Clause) GetOperator() string {
	return c.Operator
}

func (c *Clause) SetOperator(operator string) {
	c.Operator = operator
}

func (c *Clause) GetValue() any {
	return c.Value
}

func (c *Clause) SetValue(value any) {
	c.Value = value
}

func (c *Clause) GetChildren() []IClause {
	return c.Children
}

func (c *Clause) SetChildren(children []IClause) {
	c.Children = children
}
