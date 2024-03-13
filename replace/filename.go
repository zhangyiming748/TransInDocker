package replace

import (
	"log/slog"
	"regexp"
	"strings"
)

func ChinesePunctuation(str string) string {
	str = strings.Replace(str, "。", ".", -1)
	str = strings.Replace(str, "，", ",", -1)
	str = strings.Replace(str, "《", "(", -1)
	str = strings.Replace(str, "》", ")", -1)
	str = strings.Replace(str, "【", "(", -1)
	str = strings.Replace(str, "】", ")", -1)
	str = strings.Replace(str, "（", "(", -1)
	str = strings.Replace(str, "）", ")", -1)
	str = strings.Replace(str, "「", "(", -1)
	str = strings.Replace(str, "」", ")", -1)
	str = strings.Replace(str, "+", "_", -1)
	str = strings.Replace(str, "`", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "\u0000", "", -1)
	str = strings.Replace(str, "·", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	str = strings.Replace(str, "\u000D", "", -1)
	str = strings.Replace(str, "、", "", -1)
	//str = strings.Replace(str, "/", "", -1)
	str = strings.Replace(str, "！", "", -1)
	str = strings.Replace(str, "|", "", -1)
	str = strings.Replace(str, "｜", "", -1)
	str = strings.Replace(str, ":", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "&", "", -1)
	str = strings.Replace(str, "？", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "“", "", -1)
	str = strings.Replace(str, "”", "", -1)
	str = strings.Replace(str, "--", "", -1)
	str = strings.Replace(str, "_", "", -1)
	str = strings.Replace(str, "：", "", -1)
	str = strings.Replace(str, "\ufeff", "", -1)
	str = strings.Replace(str, "\n", "", 1)
	return str
}

/*
仅保留文件名中的 数字 字母 和 中文
*/
func ForFileName(name string) string {
	nStr := ""
	for _, v := range name {
		if Effective(string(v)) {
			// fmt.Printf("%d\t有效%v\n", i, string(v))
			nStr = strings.Join([]string{nStr, string(v)}, "")
		}
	}
	slog.Debug("正则表达式匹配数字字母汉字", slog.String("文件名", nStr))
	return nStr
}
func Effective(s string) bool {
	if s == " " {
		return true
	}
	num := regexp.MustCompile(`\d`)          // 匹配任意一个数字
	letter := regexp.MustCompile(`[a-zA-Z]`) // 匹配任意一个字母
	char := regexp.MustCompile(`[\p{Han}]`)  // 匹配任意一个汉字
	if num.MatchString(s) || letter.MatchString(s) || char.MatchString(s) {
		return true
	}
	return false
}

/*
替换掉失败查询留下的信息
*/
func Falied(dst string) bool {
	if strings.Contains(dst, "\u001B") {
		return true
	}
	if strings.Contains(dst, "[33mDidyoumean[1mI'malwayshereI'vefoundaverygoodprettyoutfit[22m[0m") {
		return true
	}
	if strings.Contains(dst, "[33m[WARNING]Connectiontimedout.RetryingIPv4connection.[0m") {
		return true
	}
	if strings.Contains(dst, "[33mDidyoumean[1mIt'shate,ithate,ithate,ithate.[22m[0m") {
		return true
	}
	if strings.Contains(dst, "\u001B[33m[WARNING]Connectiontimedout.RetryingIPv4connection.[0m") {
		return true
	}
	if strings.Contains(dst, "\u001B[33mDidyoumean\u001B[1mLongago,therewasasmallkingdomcalledViridian.\u001B[22m\u001B[0m") {
		return true
	}
	if strings.Contains(dst, "\u001B[33mDidyoumean\u001B[1m") {
		return true
	}
	if strings.Contains(dst, "\u001B[22m\u001B[0m") {
		return true
	}
	if strings.Contains(dst, "\u001B[22m\u001B[0m") {
		return true
	}
	return false
}
