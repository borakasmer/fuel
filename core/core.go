package core

import (
	"strings"
)

type String struct {
	Value string
}

func (str *String) Slice() string {
	if strings.Index(str.Value, ".") != -1 {
		return str.Value[0:5]
	} else {
		return str.Value[0:2]
	}
}
