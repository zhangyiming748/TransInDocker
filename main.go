package main

import (
	"fmt"
	"freeTranslate/GetFileInfo"
	"freeTranslate/replace"
	sql "freeTranslate/sql"
	"freeTranslate/translateShell"
	"freeTranslate/util"
	"io"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {
	proxy := os.Getenv("proxy")
	if proxy == "" {
		proxy = "192.168.1.5:8889"
	}
	host, port, _ := net.SplitHostPort(proxy)
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second*5); err != nil {
		slog.Error("代理网址不可达")
		os.Exit(-1)
	} else {
		slog.Info("", slog.Any("本地连接", conn.LocalAddr()), slog.Any("远程连接", conn.RemoteAddr()))
	}
	setLog()
	sql.SetEngine()
}

func main() {
	c := new(translateShell.Count)
	defer c.GetAll()
	replace.SetSensitive()
	folder := "/srt"
	//folder := "/mnt/f/srt"
	files, _ := GetFileInfo.GetAllFileInfoFast(folder, "srt")
	for _, file := range files {
		if strings.Contains(file, "origin") {
			continue
		}
		trans(file, c)
	}
}

func trans(srt string, c *translateShell.Count) {
	//host := strings.Split(util.GetVal("shell", "proxy"), ":")[0]
	//port := strings.Split(util.GetVal("shell", "proxy"), ":")[1]
	//需要一个proxy变量

	seed := rand.New(rand.NewSource(time.Now().Unix()))
	r := seed.Intn(2000)
	//中间文件名
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadByLine(srt)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		after.WriteString(fmt.Sprintf("%s\n", before[i]))
		after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]

		afterSrc := replace.GetSensitive(src)

		var dst string
		cache := new(sql.History)
		cache.Src = afterSrc
		if result := cache.FindOneBySrc(); result.Error == nil {
			dst = cache.Dst
			slog.Debug("find in cache")
			c.SetCache()
		} else {
			dst = translateShell.Translate(afterSrc, c)
			var count int
			for replace.Falied(dst) {
				if count > 3 {
					slog.Error("重试三次后依然失败", slog.String("原文", afterSrc), slog.String("译文", dst))
					dst = replace.Hans(dst)
					break
				}
				slog.Error("查询失败", slog.Int("重试", count))
				time.Sleep(1 * time.Second)
				dst = translateShell.Translate(afterSrc, c)
				count++
			}
		}
		dst = replace.GetSensitive(dst)
		// dst = replace.Hans(dst)
		slog.Info("", slog.String("文件名", tmpname), slog.String("原文", src), slog.String("译文", dst))
		after.WriteString(fmt.Sprintf("%s\n", src))
		after.WriteString(fmt.Sprintf("%s\n", dst))

		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
	}
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	exec.Command("cp", srt, origin).CombinedOutput()
	os.Rename(tmpname, srt)

}
func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	file := "/srt/trans.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	slog.SetDefault(logger)
}
