package testutil

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func ExecuteTestSuiteEntries(t *testing.T, s interface{}) {
	r := reflect.TypeOf(s)
	for i := 0; i < r.NumMethod(); i++ {
		method := r.Method(i)
		match, _ := regexp.MatchString("^Test", method.Name)
		if !match {
			continue
		}
		action := reflect.ValueOf(s).MethodByName(method.Name)
		fmt.Println("Execute :", method.Name)
		action.Call([]reflect.Value{reflect.ValueOf(t)})
	}
}
