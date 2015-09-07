package enum_util

import (
	"fmt"
	"reflect"
	"strconv"
)

type Field struct {
	Name  string
	Value int
}

func (f Field) ValueString() string {
	return strconv.Itoa(f.Value)
}

func GetFields(s interface{}) []Field {
	fields := []Field{}

	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueField := val.Field(i)
		fields = append(fields, Field{Name: typeField.Name, Value: valueField.Interface().(int)})
	}

	return fields
}

func GetFieldNames(s interface{}) []string {
	data := []string{}
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		data = append(data, typeField.Name)
	}
	return data
}

func GetFieldValue(s interface{}, value string) (int, error) {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if typeField.Name == value {
			valueField := val.Field(i)
			return valueField.Interface().(int), nil
		}
	}
	return 0, fmt.Errorf("Field Not Found (%s)", value)

}

func GetFieldName(s interface{}, value int) (string, error) {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if valueField.Interface() == value {
			typeField := val.Type().Field(i)
			return typeField.Name, nil
		}
	}
	return "", fmt.Errorf("Field Not Found (%d)", value)
}
