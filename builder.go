package sql

import (
	"github.com/Metadiv-Technology-Limited/sql/internal/clause"
	"github.com/Metadiv-Technology-Limited/sql/models"
)

/*
Eq creates a new clause with operator `=`.
*/
func Eq(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_EQ, value)
}

/*
Neq creates a new clause with operator `<>`.
*/
func Neq(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_NEQ, value)
}

/*
Gt creates a new clause with operator `>`.
*/
func Gt(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_GT, value)
}

/*
Gte creates a new clause with operator `>=`.
*/
func Gte(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_GTE, value)
}

/*
Lt creates a new clause with operator `<`.
*/
func Lt(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_LT, value)
}

/*
Lte creates a new clause with operator `<=`.
*/
func Lte(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_LTE, value)
}

/*
Like creates a new clause with operator `LIKE`.
*/
func Like(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_LIKE, value)
}

/*
NotLike creates a new clause with operator `NOT LIKE`.
*/
func NotLike(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_NOT_LIKE, value)
}

/*
Similar creates a new clause with operator "LIKE" and the value "%value%".
*/
func Similar(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_LIKE, "%"+value.(string)+"%")
}

/*
NotSimilar creates a new clause with operator "NOT LIKE" and the value "%value%".
*/
func NotSimilar(field string, value any) models.IClause {
	return clause.NewClause(field, clause.OP_NOT_LIKE, "%"+value.(string)+"%")
}

/*
In creates a new clause with operator `IN`.
*/
func In(field string, value ...any) models.IClause {
	return clause.NewClause(field, clause.OP_IN, value)
}

/*
NotIn creates a new clause with operator `NOT IN`.
*/
func NotIn(field string, value ...any) models.IClause {
	return clause.NewClause(field, clause.OP_NOT_IN, value)
}

/*
IsNull creates a new clause with operator `IS NULL`.
*/
func IsNull(field string) models.IClause {
	return clause.NewClause(field, clause.OP_IS_NULL, nil)
}

/*
IsNotNull creates a new clause with operator `IS NOT NULL`.
*/
func IsNotNull(field string) models.IClause {
	return clause.NewClause(field, clause.OP_NOT_NULL, nil)
}

/*
And creates a new clause with operator `AND`.
*/
func And(children ...models.IClause) models.IClause {
	return clause.NewClause("", clause.OP_AND, nil, children...)
}

/*
Or creates a new clause with operator `OR`.
*/
func Or(children ...models.IClause) models.IClause {
	return clause.NewClause("", clause.OP_OR, nil, children...)
}

/*
Between creates a new clause that checks if a field is between two values.
*/
func Between(field string, from, to any, includeEdgeTo ...bool) models.IClause {
	if len(includeEdgeTo) > 0 && includeEdgeTo[0] {
		return And(Gte(field, from), Lte(field, to))
	}
	return And(Gt(field, from), Lt(field, to))
}

/*
NotBetween creates a new clause that checks if a field is not between two values.
*/
func NotBetween(field string, from, to any, includeEdgeTo ...bool) models.IClause {
	if len(includeEdgeTo) > 0 && includeEdgeTo[0] {
		return Or(Lt(field, from), Gt(field, to))
	}
	return Or(Lte(field, from), Gte(field, to))
}
