package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func ReadByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		log.Printf("按行读文件出错:%v\n", err)
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

// 按行写文件
func WriteByLine(fp string, s []string) {
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
	return

}
func IsExist(folderPath string) bool {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在")
		return false
	} else {
		fmt.Println("文件夹存在")
		return true
	}
}
func IsExistCmd(cmds ...string) bool {
	for _, cmd := range cmds {
		//cmd := "ls" // 需要测试的命令
		_, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Printf("命令 %s 不存在\n", cmd)
			return false
		} else {
			fmt.Printf("命令 %s 存在\n", cmd)
		}
	}
	return true
}
