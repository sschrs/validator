package govalidator

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	Name  string
	Value reflect.Value
	Tag   string
}

func Validate(s interface{}) (bool, string) {
	fields := getFieldsFromStruct(s)
	for _, field := range fields {
		ok, message := validateField(field)
		if !ok {
			return ok, message
		}
	}
	return true, ""
}

func validateField(field Field) (bool, string) {
	rules := strings.Split(field.Tag, defaultTagSeperator)
	for _, rule := range rules {
		if strings.Contains(rule, "=") {
			parseRule := strings.Split(rule, "=")
			if mapRulesToFuncs[parseRule[0]] == nil {
				panic(fmt.Sprintf("unknown validation tag: %s for %s", rule, field.Name))
			}
			ok, message := mapRulesToFuncs[parseRule[0]](field, parseRule[1])
			if !ok {
				return ok, message
			}
		} else {
			if mapRulesToFuncs[rule] == nil {
				panic(fmt.Sprintf("unknown validation tag: %s for %s", rule, field.Name))
			}
			ok, message := mapRulesToFuncs[rule](field, "")
			if !ok {
				return ok, message
			}
		}
	}
	return true, ""
}

func getFieldsFromStruct(s interface{}) []Field {
	var fields []Field
	fieldCount := reflect.ValueOf(s).Elem().NumField()

	for i := 0; i < fieldCount; i++ {
		field := reflect.TypeOf(s).Elem().Field(i)
		tag := field.Tag.Get(defaultTagName)
		fieldValue := reflect.ValueOf(s).Elem().Field(i)
		newField := Field{
			Name:  field.Name,
			Value: fieldValue,
			Tag:   tag,
		}
		fields = append(fields, newField)
	}
	return fields
}

func AddCustomValidation(validationName string, validationFunc func(field Field, value string) (bool, string)) {
	mapRulesToFuncs[validationName] = validationFunc
}
