package cast

// 忘れるので集めとく

import (
	"strconv"
)

// 数字を文字列に
func Int2Str(i int) string {
	return strconv.Itoa(i)
}

// [文字列]文字列を、[文字列]インターフェースに
func StrStr2StrInt(tmp map[string]string) map[string]interface{} {
	args := map[string]interface{}{}
	for k, v := range tmp {
		args[k] = v
	}
	return args
}

// 文字列を数字に
func Str2Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
