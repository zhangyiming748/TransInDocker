package constant

import (
	"log"
	"os"
	"strings"
)

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
