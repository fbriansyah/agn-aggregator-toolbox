package util

import "database/sql"

func ValidNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		Valid:  true,
		String: str,
	}
}
