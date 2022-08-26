package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshallJSON() ([]byte, error) {
	value := fmt.Sprintf("%d mins", r)

	quoted := strconv.Quote(value)
	return []byte(quoted), nil
}
