package constant

import (
	"log"
	"os"
	"strings"
)

var (
	BitRate = map[string]string{
		"avc":  "5000K",
		"hevc": "1800K",
	}
)

const (
	Type      = iota + 1
	Kilobyte  = 1000 * Type
	Megabyte  = 1000 * Kilobyte
	Gigabyte  = 1000 * Megabyte
	Terabyte  = 1000 * Gigabyte
	Petabyte  = 1000 * Terabyte
	Exabyte   = 1000 * Petabyte
	Zettabyte = 1000 * Exabyte
	Yottabyte = 1000 * Zettabyte
)
const HTTPS = "https://fanyi-api.baidu.com/api/trans/vip/translate"
const (
	JP = "jp"
	EN = "en"
	ZH = "zh"
)

var T2B = map[string]string{
	"en":    "en",
	"ja":    "jp",
	"zh-CN": "zh",
	"ko":    "kor", // 韩语
	"th":    "th",  // 泰语
	"de":    "de",  //德语
	"fr":    "fra", //法语
	"ru":    "ru",  // 俄语
	"sp":    "spa", // 西班牙语
}
var (
	root  = "/Users/zen/Github/FastYt-dlp"
	proxy = "192.168.1.20:8889"
	from  = ":zh-CN"
)

func SetRoot(s string) {
	root = s
}

func GetRoot() string {
	return root
}

func SetProxy() {
	env := os.Getenv("proxy")
	if env != "" {
		proxy = env
		log.Printf("读取到proxy环境变量:%v\n", proxy)
	} else {
		log.Printf("未读取到proxy环境变量:%v\n", proxy)
	}
}
func GetProxy() string {
	return proxy
}
func SetFrom() {
	env := os.Getenv("from")
	if env != "" {
		from = strings.Join([]string{env, from}, "")
		log.Printf("读取到from环境变量:%v\n", from)
	} else {
		log.Printf("未读取到from环境变量:%v\n", from)
	}
}
func GetFrom() string {
	return from
}
