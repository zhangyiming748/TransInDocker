package main

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"zhangyiming748/TransInDocker/constant"
	"zhangyiming748/TransInDocker/replace"
	sql "zhangyiming748/TransInDocker/sql"
	"zhangyiming748/TransInDocker/translateShell"
	"zhangyiming748/TransInDocker/util"
)

func init() {
	if !util.IsExistCmd("trans") {
		log.Fatalln("缺少运行程序的基础命令")
	}
	proxy := constant.GetProxy()
	host, port, _ := net.SplitHostPort(proxy)
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second*5); err != nil {
		log.Fatalln("代理网址不可达")
	} else {
		log.Printf("本地连接%v\t远程连接%v\n", conn.LocalAddr(), conn.RemoteAddr())
	}
	setLog()
	sql.Initial()
}

var (
	db = sql.GetDatabase()
)

func main() {
	c := new(translateShell.Count)
	defer c.GetAll()
	replace.SetSensitive()
	if r := os.Getenv("root"); r == "" {
		log.Printf("没有设置root变量,使用默认:%v\n", constant.GetRoot())
	} else {
		constant.SetRoot(r)
	}
	files, _ := getFilesWithExt(constant.GetRoot(), ".srt")
	for _, file := range files {
		if strings.Contains(file, "origin") {
			continue
		}
		trans(file, c)
	}
}

func trans(srt string, c *translateShell.Count) {
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

		if get, err := db.Hash().Get("translations", src); err == nil {
			dst = get.String()
			fmt.Println("find in cache")
			c.SetCache()
		} else {
			dst = translateShell.Translate(afterSrc, c)
			var count int
			for replace.Falied(dst) {
				if count > 3 {
					log.Printf("重试三次后依然失败srt=%v\tdst=%v\n", afterSrc, dst)
					dst = replace.Hans(dst)
					break
				}
				log.Printf("查询失败\t重试%v\n", count)
				time.Sleep(1 * time.Second)
				dst = translateShell.Translate(afterSrc, c)
				count++
			}
		}
		dst = replace.GetSensitive(dst)
		db.Hash().Set("translations", src, dst)
		fmt.Printf("文件名:%v\t原文:%v\t译文:%v\n", tmpname, src, dst)
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
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{constant.GetRoot(), "mylog.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}

func getFilesWithExt(path string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			files = append(files, absPath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
