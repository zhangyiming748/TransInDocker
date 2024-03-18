package translateShell

import (
	"freeTranslate/replace"
	"freeTranslate/sql"
	"log/slog"
	"os"
	"time"
)

const (
	TIMEOUT = 8 //second
)

func Translate(src string, c *Count) string {
	//trans -brief ja:zh "私の手の動きに合わせて、そう"
	his := new(sql.History)
	defer func() {
		his.SetOne()
	}()
	bing := make(chan string, 1)
	google := make(chan string, 1)

	// from := os.Getenv("from")
	// to := os.Getenv("to")
	proxy := os.Getenv("proxy")
	if proxy == "" {
		proxy = "192.168.1.5:8889"
	}
	//language := strings.Join([]string{from, to}, ":")
	language := ":zh-CN"

	go TransByGoogle(proxy, language, src, google)
	go TransByBing(proxy, language, src, bing)

	var dst string
	select {
	case dst = <-bing:
		his.Source = "bing"
		c.SetBing()
	case dst = <-google:
		his.Source = "google"
		c.SetGoogle()
	case <-time.After(TIMEOUT * time.Second):
		slog.Error("单词翻译出现严重问题")
	}

	dst = replace.ChinesePunctuation(dst)
	dst = replace.Hans(dst)
	his.From = "auto"
	his.To = "zh-CN"
	his.Src = src
	his.Dst = dst

	return dst
}
