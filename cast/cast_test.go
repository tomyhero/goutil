package cast

import (
	"testing"
)

func TestInt2Str(t *testing.T) {
	if Int2Str(10) != "10" {
		t.Fail()
	}
}

func TestStrStr2StrInt(t *testing.T) {
	d := StrStr2StrInt(map[string]string{"a": "b"})
	if d["a"].(string) != "b" {
		t.Fail()
	}
}

func TestStr2Int(t *testing.T) {
	if Str2Int("10") != 10 {
		t.Fail()
	}
}
