package govalidator

import (
	"fmt"
	"govalidator/messages"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

var mapRulesToFuncs = map[string]func(field Field, value string) (bool, string){
	"notNull":              notNull,
	"email":                email,
	"numeric":              numeric,
	"containsString":       containsString,
	"containsNumber":       containsNumber,
	"containsSpecialChars": containsSpecialChars,
	"eq":                   eq,
	"gt":                   gt,
	"lt":                   lt,
	"lte":                  lte,
	"gte":                  gte,
	"ne":                   ne,
	"contains":             contains,
	"notContains":          notContains,
	"begin":                begin,
	"notBegin":             notBegin,
	"end":                  end,
	"notEnd":               notEnd,
}

func notNull(field Field, value string) (bool, string) {
	var isNil bool
	k := field.Value.Kind().String()
	switch k {
	case "chan", "func", "map":
		isNil = field.Value.IsNil()
	default:
		isNil = false
	}

	if field.Value.String() == "" || isNil {
		return false, fmt.Sprintf("%s %s", field.Name, messages.Languages[defaultLanguage]["notNull"])
	}
	return true, ""
}

func email(field Field, value string) (bool, string) {
	emailAddress := field.Value.String()
	_, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return false, fmt.Sprintf("%s %s", field.Name, messages.Languages[defaultLanguage]["email"])
	}
	return true, ""
}

func numeric(field Field, value string) (bool, string) {
	ok := field.Value.CanInt()
	if !ok && field.Value.String() != "" {
		return false, fmt.Sprintf("%s %s", field.Name, messages.Languages[defaultLanguage]["numeric"])
	}
	return true, ""
}

func containsString(field Field, value string) (bool, string) {
	fieldValue := field.Value.String()
	if fieldValue != "" {
		var check = regexp.MustCompile(".*[a-zA-Z].*").MatchString
		ok := check(fieldValue)
		if !ok {
			return false, fmt.Sprintf("%s %s", field.Name, messages.Languages[defaultLanguage]["containsString"])
		}
	}
	return true, ""
}

func containsNumber(field Field, value string) (bool, string) {
	fieldValue := field.Value.String()
	if fieldValue != "" {
		var check = regexp.MustCompile(".*\\d+.*").MatchString
		ok := check(fieldValue)
		if !ok {
			return false, fmt.Sprintf("%s %s", field.Name, messages.Languages[defaultLanguage]["containsNumber"])
		}
	}
	return true, ""
}

func containsSpecialChars(field Field, value string) (bool, string) {
	fieldValue := field.Value.String()
	if fieldValue != "" {
		var check = regexp.MustCompile(".*[" + specialChars + "]").MatchString
		ok := check(fieldValue)
		if !ok {
			return false, fmt.Sprintf("%s %s", field.Name, messages.Languages[defaultLanguage]["containsSpecialChars"])
		}
	}
	return true, ""
}

func eq(field Field, value string) (bool, string) {
	_value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("mismatched types. value should be an integer: %s for %s", field.Tag, field.Name))
	}

	switch field.Value.Kind().String() {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		if _value != field.Value.Int() {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["eqNumber"], _value)
		}

	case "string":
		if _value != int64(field.Value.Len()) {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["eqString"], _value)
		}
	}
	return true, ""
}

func gt(field Field, value string) (bool, string) {
	_value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("mismatched types. value should be an integer: %s for %s", field.Tag, field.Name))
	}

	switch field.Value.Kind().String() {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		if _value >= field.Value.Int() {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["gtNumber"], _value)
		}

	case "string":
		if _value >= int64(field.Value.Len()) {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["gtString"], _value)
		}
	}
	return true, ""
}

func lt(field Field, value string) (bool, string) {
	_value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("mismatched types. value should be an integer: %s for %s", field.Tag, field.Name))
	}

	switch field.Value.Kind().String() {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		if _value <= field.Value.Int() {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["ltNumber"], _value)
		}

	case "string":
		if _value <= int64(field.Value.Len()) {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["ltString"], _value)
		}
	}
	return true, ""
}

func lte(field Field, value string) (bool, string) {
	_value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("mismatched types. value should be an integer: %s for %s", field.Tag, field.Name))
	}

	switch field.Value.Kind().String() {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		if _value < field.Value.Int() {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["lteNumber"], _value)
		}

	case "string":
		if _value < int64(field.Value.Len()) {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["lteString"], _value)
		}
	}
	return true, ""
}

func gte(field Field, value string) (bool, string) {
	_value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("mismatched types. value should be an integer: %s for %s", field.Tag, field.Name))
	}

	switch field.Value.Kind().String() {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		if _value > field.Value.Int() {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["gteNumber"], _value)
		}

	case "string":
		if _value > int64(field.Value.Len()) {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["gteString"], _value)
		}
	}
	return true, ""
}

func ne(field Field, value string) (bool, string) {
	_value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("mismatched types. value should be an integer: %s for %s", field.Tag, field.Name))
	}

	switch field.Value.Kind().String() {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		if _value == field.Value.Int() {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["neNumber"], _value)
		}

	case "string":
		if _value == int64(field.Value.Len()) {
			return false, fmt.Sprintf("%s %s %d", field.Name, messages.Languages[defaultLanguage]["neString"], _value)
		}
	}
	return true, ""
}

func contains(field Field, value string) (bool, string) {
	if !strings.Contains(field.Value.String(), value) {
		return false, fmt.Sprintf("%s %s '%s'", field.Name, messages.Languages[defaultLanguage]["contains"], value)
	}
	return true, ""
}

func notContains(field Field, value string) (bool, string) {
	if strings.Contains(field.Value.String(), value) {
		return false, fmt.Sprintf("%s %s '%s'", field.Name, messages.Languages[defaultLanguage]["notContains"], value)
	}
	return true, ""
}

func begin(field Field, value string) (bool, string) {
	if !strings.HasPrefix(field.Value.String(), value) {
		return false, fmt.Sprintf("%s %s '%s'", field.Name, messages.Languages[defaultLanguage]["begin"], value)
	}
	return true, ""
}

func notBegin(field Field, value string) (bool, string) {
	if strings.HasPrefix(field.Value.String(), value) {
		return false, fmt.Sprintf("%s %s '%s'", field.Name, messages.Languages[defaultLanguage]["notBegin"], value)
	}
	return true, ""
}

func end(field Field, value string) (bool, string) {
	if !strings.HasSuffix(field.Value.String(), value) {
		return false, fmt.Sprintf("%s %s '%s'", field.Name, messages.Languages[defaultLanguage]["end"], value)
	}
	return true, ""
}

func notEnd(field Field, value string) (bool, string) {
	if strings.HasSuffix(field.Value.String(), value) {
		return false, fmt.Sprintf("%s %s '%s'", field.Name, messages.Languages[defaultLanguage]["notEnd"], value)
	}
	return true, ""
}
