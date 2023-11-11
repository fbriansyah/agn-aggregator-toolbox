package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidNullString(t *testing.T) {
	str := "Test"
	res := ValidNullString(str)

	require.Equal(t, str, res.String)
	require.Equal(t, true, res.Valid)
}

func TestValidNullInt32(t *testing.T) {
	n1 := int32(32)
	n2 := ValidNullInt32(n1)

	require.Equal(t, n1, n2.Int32)
	require.Equal(t, true, n2.Valid)
}

func TestLikeCondition(t *testing.T) {
	field := "field"
	value := "value"
	expected := field + " LIKE '%" + value + "%'"

	var qw QueryWhere
	qw.Like(field, value)

	require.Equal(t, false, qw.IsEmpty())
	require.Equal(t, expected, qw.Build())
}

func TestEqual(t *testing.T) {
	field := "field"
	value := "value"
	expected := fmt.Sprintf("%s = '%v'", field, value)

	var qw QueryWhere
	qw.Equal(field, value)

	require.Equal(t, false, qw.IsEmpty())
	require.Equal(t, expected, qw.Build())
}

func TestGreater(t *testing.T) {
	field := "field"
	value := 50
	expected := fmt.Sprintf("%s > %v", field, value)

	var qw QueryWhere
	qw.Greater(field, value)

	require.Equal(t, false, qw.IsEmpty())
	require.Equal(t, expected, qw.Build())
}

func TestLess(t *testing.T) {
	field := "field"
	value := 50
	expected := fmt.Sprintf("%s < %v", field, value)

	var qw QueryWhere
	qw.Less(field, value)

	require.Equal(t, false, qw.IsEmpty())
	require.Equal(t, expected, qw.Build())
}

func TestDate(t *testing.T) {
	field := "field"
	value := "2023-11-10"
	expected := fmt.Sprintf("DATE(%s) = '%s'", field, value)

	var qw QueryWhere
	qw.Date(field, value)

	require.Equal(t, false, qw.IsEmpty())
	require.Equal(t, expected, qw.Build())
}

func TestIsEmpty(t *testing.T) {
	var qw QueryWhere

	require.True(t, qw.IsEmpty())
}

func TestBuild(t *testing.T) {
	expected := "field1 LIKE '%value1%' AND DATE(field2) = 'value2'"

	var qw QueryWhere
	qw.Like("field1", "value1")
	qw.Date("field2", "value2")

	require.False(t, qw.IsEmpty())
	require.Equal(t, expected, qw.Build())
}
