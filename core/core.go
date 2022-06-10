package core

import (
	"strings"
)

type String struct {
	Value string
}

func (str *String) Slice() string {
	// prevent panicking
	if len(str.Value) == 0 {
		return ""
	}
	if strings.Contains(str.Value, ".") {
		return str.Value[0:5]
	}
	return str.Value[0:2]
}
