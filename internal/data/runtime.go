package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

const minsSuffix = "mins"

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshallJSON() ([]byte, error) {
	value := fmt.Sprintf("%d mins", r)

	quoted := strconv.Quote(value)
	return []byte(quoted), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquoted, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquoted, " ")

	if len(parts) != 2 || parts[1] != minsSuffix {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
