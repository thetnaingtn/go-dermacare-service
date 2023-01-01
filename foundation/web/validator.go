package web

import (
	"reflect"
	"strings"
)

func ValidatorTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}

	return name
}
