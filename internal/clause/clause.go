package clause

import (
	"strings"

	"github.com/Metadiv-Technology-Limited/sql/models"
)

/*
NewClause creates a new clause.
*/
func NewClause(field, operator string, value interface{}, children ...models.IClause) models.IClause {
	return &models.Clause{
		Field:    safeField(field),
		Operator: operator,
		Value:    value,
		Children: children,
	}
}

/*
safeField is a utility function that makes a field safe.
We append ` to the field name to avoid conflict with SQL keywords.
*/
func safeField(field string) string {
	field = strings.ReplaceAll(field, "`", "")
	field = strings.ReplaceAll(field, ".", "`.`")
	return "`" + field + "`"
}
