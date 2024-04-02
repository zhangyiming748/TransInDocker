package replace

import (
	"bufio"
	"fmt"
	"freeTranslate/sql"
	"io"
	"log/slog"
	"os"
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
	lines := readByLine("/srt/sensitive.txt")
	for _, line := range lines {
		before := strings.Split(line, ":")[0]
		after := strings.Split(line, ":")[1]
		slog.Info("写入敏感词", slog.String("before", before), slog.String("after", after))
		Sensitive[before] = after
	}
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
	done = remove331x220(done)

	return done
}

/*
golang 实现 从字符串中找到以331开头220结尾的子字符串 删除后返回新的字符串
*/
func remove331x220(s string) string {
	re := regexp.MustCompile(`331.*?220`)
	return re.ReplaceAllString(s, "")
}
func readByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		slog.Warn("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}
