package sql

/*
Eq creates a new clause with operator `=`.
*/
func Eq(field string, value any) *Clause {
	return newClause(field, OP_EQ, value)
}

/*
Neq creates a new clause with operator `<>`.
*/
func Neq(field string, value any) *Clause {
	return newClause(field, OP_NEQ, value)
}

/*
Gt creates a new clause with operator `>`.
*/
func Gt(field string, value any) *Clause {
	return newClause(field, OP_GT, value)
}

/*
Gte creates a new clause with operator `>=`.
*/
func Gte(field string, value any) *Clause {
	return newClause(field, OP_GTE, value)
}

/*
Lt creates a new clause with operator `<`.
*/
func Lt(field string, value any) *Clause {
	return newClause(field, OP_LT, value)
}

/*
Lte creates a new clause with operator `<=`.
*/
func Lte(field string, value any) *Clause {
	return newClause(field, OP_LTE, value)
}

/*
Like creates a new clause with operator `LIKE`.
*/
func Like(field string, value any) *Clause {
	return newClause(field, OP_LIKE, value)
}

/*
NotLike creates a new clause with operator `NOT LIKE`.
*/
func NotLike(field string, value any) *Clause {
	return newClause(field, OP_NOT_LIKE, value)
}

/*
Similar creates a new clause with operator "LIKE" and the value "%value%".
*/
func Similar(field string, value any) *Clause {
	return newClause(field, OP_LIKE, "%"+value.(string)+"%")
}

/*
NotSimilar creates a new clause with operator "NOT LIKE" and the value "%value%".
*/
func NotSimilar(field string, value any) *Clause {
	return newClause(field, OP_NOT_LIKE, "%"+value.(string)+"%")
}

/*
In creates a new clause with operator `IN`.
*/
func In(field string, value ...any) *Clause {
	return newClause(field, OP_IN, value)
}

/*
NotIn creates a new clause with operator `NOT IN`.
*/
func NotIn(field string, value ...any) *Clause {
	return newClause(field, OP_NOT_IN, value)
}

/*
IsNull creates a new clause with operator `IS NULL`.
*/
func IsNull(field string) *Clause {
	return newClause(field, OP_IS_NULL, nil)
}

/*
IsNotNull creates a new clause with operator `IS NOT NULL`.
*/
func IsNotNull(field string) *Clause {
	return newClause(field, OP_NOT_NULL, nil)
}

/*
And creates a new clause with operator `AND`.
*/
func And(children ...*Clause) *Clause {
	return newClause("", OP_AND, nil, children...)
}

/*
Or creates a new clause with operator `OR`.
*/
func Or(children ...*Clause) *Clause {
	return newClause("", OP_OR, nil, children...)
}

/*
Between creates a new clause that checks if a field is between two values.
*/
func Between(field string, from, to any, includeEdgeTo ...bool) *Clause {
	if len(includeEdgeTo) > 0 && includeEdgeTo[0] {
		return And(Gte(field, from), Lte(field, to))
	}
	return And(Gt(field, from), Lt(field, to))
}

/*
NotBetween creates a new clause that checks if a field is not between two values.
*/
func NotBetween(field string, from, to any, includeEdgeTo ...bool) *Clause {
	if len(includeEdgeTo) > 0 && includeEdgeTo[0] {
		return Or(Lt(field, from), Gt(field, to))
	}
	return Or(Lte(field, from), Gte(field, to))
}
