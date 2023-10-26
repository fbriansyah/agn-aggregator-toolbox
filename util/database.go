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

type QueryWhere struct {
	Wheres []string
}

func (q *QueryWhere) Like(field string, value string) {
	q.Wheres = append(q.Wheres, field+" LIKE '%"+value+"%'")
}

func (q *QueryWhere) Equal(field string, value any) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("%s = '%v'", field, value))
}

func (q *QueryWhere) Greater(field string, value any) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("%s > %v", field, value))
}

func (q *QueryWhere) Less(field string, value any) {
	q.Wheres = append(q.Wheres, fmt.Sprintf("%s < %v", field, value))
}

func (q *QueryWhere) IsEmpty() bool {
	return len(q.Wheres) < 1
}

func (q *QueryWhere) Build() string {
	return strings.Join(q.Wheres, " AND ")
}
