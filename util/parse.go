package util

import "strconv"

// StrToInt32 parse string to int32
func StrToInt32(str string) (int32, error) {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1, err
	}

	return int32(i), nil
}
