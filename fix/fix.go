package fix

import (
	"fmt"
	"regexp"
)

func filterString(s string) string {
	reg := regexp.MustCompile(`[0-9a-zA-Z\s\p{Han}]+`)

	// 提取匹配的结果
	result := reg.FindAllString(s, -1)

	// 输出结果
	return fmt.Sprintln(result)

}

func Fix33m(str string) string {
	//str := "This is a test string with [33mDidyoumean"

	// 使用正则表达式匹配 "[33mDidyoumean" 字样
	reg := regexp.MustCompile(`\[33m`)

	// 判断字符串中是否包含 "[33mDidyoumean" 字样
	if reg.MatchString(str) {
		reg := regexp.MustCompile(`[^\p{Han}]+`)
		//reg := regexp.MustCompile(`[^\x00-\x7F]+`)

		// 替换所有匹配结果为空字符串
		newStr := reg.ReplaceAllString(str, "")

		// 输出结果
		return newStr
	} else {
		return str
	}
}
