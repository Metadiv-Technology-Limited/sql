package clause

import (
	"strings"

	"github.com/Metadiv-Technology-Limited/sql/models"

	"gorm.io/gorm"
)

/*
Consume consumes a clause and returns *gorm.DB
*/
func Consume(tx *gorm.DB, cls models.IClause) *gorm.DB {
	if cls == nil {
		return tx
	}
	stm, values := build(cls, make([]any, 0))
	return tx.Where(stm, values...)
}

/*
build is a utility function that builds a clause.
*/
func build(cls models.IClause, values []any) (string, []any) {
	var stm string
	switch strings.ToUpper(cls.GetOperator()) {
	case OP_AND, OP_OR:
		stm, values = buildHasChildren(cls, values)
	case OP_IS_NULL, OP_NOT_NULL:
		stm = cls.GetField() + " " + cls.GetOperator()
	case OP_IN, OP_NOT_IN:
		stm = cls.GetField() + " " + cls.GetOperator() + " (" + strings.TrimRight(strings.Repeat("?,", len(cls.GetValue().([]interface{}))), ",") + ")"
		values = append(values, cls.GetValue().([]interface{})...)
	default:
		stm = cls.GetField() + " " + cls.GetOperator() + " ?"
		values = append(values, cls.GetValue())
	}
	return stm, values
}

/*
buildHasChildren is a utility function that builds a clause with children.
We no need to export this function.
*/
func buildHasChildren(cls models.IClause, values []interface{}) (string, []interface{}) {
	var buf []string
	for _, child := range cls.GetChildren() {
		s, v := build(child, values)
		buf = append(buf, s)
		values = v
	}
	stm := "(" + strings.Join(buf, " "+cls.GetOperator()+" ") + ")"
	return stm, values
}
