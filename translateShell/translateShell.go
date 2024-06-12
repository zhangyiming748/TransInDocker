package translateShell

import (
	"log"
	"sync"
	"time"
	"zhangyiming748/TransInDocker/constant"
	"zhangyiming748/TransInDocker/replace"
)

const (
	TIMEOUT = 8 //second
)

type Result struct {
	From string // 来源
	Dst  string // 翻译内容
}

func (r *Result) SetResult(from, dst string) {
	r.From = from
	r.Dst = dst
}
func (r Result) GetResult() Result {
	return r
}
func Translate(src string, c *Count) (dst string) {
	//trans -brief ja:zh "私の手の動きに合わせて、そう"
	ch := make(chan Result)
	var once sync.Once
	proxy := constant.GetProxy()
	language := constant.GetFrom()
	retry := 0
	for {
		go TransByGoogle(proxy, language, src, ch, c, &once)
		go TransByBing(proxy, language, src, ch, c, &once)
		//使用同一个通道 传递结构体 标明来源
		var result Result
		select {
		case result = <-ch:
			if result.From == "google" {
				c.SetGoogle()
			} else if result.From == "bing" {
				c.SetBing()
			}
			dst = result.Dst
		case <-time.After(TIMEOUT * time.Second):
			log.Println("trans超时,使用deeplx翻译")
			dst = DeepXl(src)
			log.Printf("deepXL翻译结果:%v\n", dst)
			c.SetDeeplx()
		}
		if dst != "" {
			break
		} else {
			retry++
			log.Printf("查询结果为空retry:%v\n", retry)
		}
		if retry >= 3 {
			break
		}
	}
	dst = replace.ChinesePunctuation(dst)
	dst = replace.Hans(dst)
	return dst
}
