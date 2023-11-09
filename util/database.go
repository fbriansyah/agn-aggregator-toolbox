package util

import (
	"database/sql"
	"fmt"
	"strings"
)

func ValidNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		Valid:  true,
		String: str,
	}
}

func ValidNullInt32(n int32) sql.NullInt32 {
	return sql.NullInt32{
		Valid: true,
		Int32: n,
	}
}

type QueryWhere struct {
	Wheres []string
}

// Like create like condition (e.g. field LIKE '%value%').
func (q *QueryWhere) Like(field string, value string) {
	q.Wheres = append(q.Wheres, field+" LIKE '%"+value+"%'")
}

// Equal create equality condition (e.g. field = 'value').
func (q *QueryWhere) Equal(field string, value any) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("%s = '%v'", field, value))
}

// Greater create greater check (e.g. field > value)
func (q *QueryWhere) Greater(field string, value any) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("%s > %v", field, value))
}

// Lest create less than check (e.g. field < value)
func (q *QueryWhere) Less(field string, value any) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("%s < %v", field, value))
}

func (q *QueryWhere) Date(field string, value string) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("DATE(%s) = '%s'", field, value))
}

// IsEmpty check if the query container empty or not
func (q *QueryWhere) IsEmpty() bool {
	return len(q.Wheres) < 1
}

// Build building query, join the where array with ' AND '.
func (q *QueryWhere) Build() string {
	return strings.Join(q.Wheres, " AND ")
}
