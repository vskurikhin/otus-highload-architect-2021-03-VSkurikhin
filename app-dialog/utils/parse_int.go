package utils

import "strconv"

const (
	PARSE_INT_BASE     = 10
	PARSE_INT_BIT_SIZE = 32
)

func ParseInt(i string) (int, error) {

	result, err := strconv.ParseInt(i, PARSE_INT_BASE, PARSE_INT_BIT_SIZE)

	if err != nil {
		return 0, err
	}
	return int(result), nil
}
