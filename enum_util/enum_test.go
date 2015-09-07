package enum_util

import (
	"reflect"
	"testing"
)

type TEST_ENUM_ENTRY struct {
	OK   int
	FAIL int
	NG   int
}

func TestGetFields(t *testing.T) {
	entry := &TEST_ENUM_ENTRY{OK: 1, FAIL: 2, NG: 3}
	fields := GetFields(entry)

	if fields[0].Name != "OK" {
		t.Error("Fail to get Name")
	}

	if fields[0].Value != 1 {
		t.Error("Fail to get Value")
	}

}

func TestGetFieldNames(t *testing.T) {
	entry := &TEST_ENUM_ENTRY{OK: 1, FAIL: 2, NG: 3}
	names := GetFieldNames(entry)

	if !reflect.DeepEqual(names, []string{"OK", "FAIL", "NG"}) {
		t.Error("not match")
	}

}

func TestGetValue(t *testing.T) {
	entry := &TEST_ENUM_ENTRY{OK: 1, FAIL: 2, NG: 3}

	value, err := GetFieldValue(entry, "OK")

	if err != nil {
		t.Error(err)
	}

	if value != 1 {
		t.Error("error", value)
	}

}

func TestGetFieldName(t *testing.T) {
	entry := &TEST_ENUM_ENTRY{OK: 1, FAIL: 2, NG: 3}

	name, err := GetFieldName(entry, 1)

	if err != nil {
		t.Error(err)
	}

	if name != "OK" {
		t.Error("error", name)
	}

	name, err = GetFieldName(entry, 2)

	if err != nil {
		t.Error(err)
	}

	if name != "FAIL" {
		t.Error("error", name)
	}
}
