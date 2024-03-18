package replace

import (
	"freeTranslate/sql"
	"log/slog"
	"regexp"
	"strings"
)

var Sensitive = map[string]string{}

func GetSensitive(str string) string {
	for k, v := range Sensitive {
		if strings.Contains(str, k) {
			strings.Replace(str, k, v, -1)
			slog.Debug("替换生效")
		}
	}
	return str
}
func SetSensitive() {
	m := new(sql.Sensitive)
	ss := m.GetAll()
	for _, s := range ss {
		slog.Info("加载敏感词", slog.String("before", s.Src), slog.String("after", s.Dst))
		Sensitive[s.Src] = s.Dst
	}
}

/*
所有符号替换为空格
*/
func space(str string) string {
	str = strings.Replace(str, "。", " ", -1)
	str = strings.Replace(str, "，", " ", -1)
	str = strings.Replace(str, "《", " ", -1)
	str = strings.Replace(str, "》", " ", -1)
	str = strings.Replace(str, "【", " ", -1)
	str = strings.Replace(str, "】", " ", -1)
	str = strings.Replace(str, "（", " ", -1)
	str = strings.Replace(str, "）", " ", -1)
	str = strings.Replace(str, "「", " ", -1)
	str = strings.Replace(str, "」", " ", -1)
	str = strings.Replace(str, "+", " ", -1)
	str = strings.Replace(str, ".", " ", 1)
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Replace(str, "`", " ", -1)
	str = strings.Replace(str, "·", " ", -1)
	str = strings.Replace(str, "、", " ", -1)
	str = strings.Replace(str, "！", " ", -1)
	str = strings.Replace(str, "|", " ", -1)
	str = strings.Replace(str, "｜", " ", -1)
	str = strings.Replace(str, ":", " ", -1)
	str = strings.Replace(str, " ", " ", -1)
	str = strings.Replace(str, "&", " ", -1)
	str = strings.Replace(str, "？", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Replace(str, " ", " ", -1)
	str = strings.Replace(str, "“", " ", -1)
	str = strings.Replace(str, "”", " ", -1)
	str = strings.Replace(str, "--", " ", -1)
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Replace(str, "：", " ", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
func Hans(input string) string {
	//input := "Hello, 你好！123abc"
	input = space(input)
	done := ""
	//reg := regexp.MustCompile(`\p{Han}|\d|[a-zA-Z]|\s`)
	reg := regexp.MustCompile(`\p{Han}|\d|\s`)
	matches := reg.FindAllString(input, -1)
	for _, match := range matches {
		//fmt.Printf("%d,%s", i, match)
		done = strings.Join([]string{done, match}, "")
	}
	return done
}
