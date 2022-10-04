package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func tailflag() (int, string) {
	var line int
	var path string
	var help bool
	flag.IntVar(&line, "f", 0, "读取文件最后多少字节, 如果不指定,则从文件开头读起")
	flag.StringVar(&path, "p", "", "文件路径")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("usage: tailf -f [byte] -p path")
		flag.PrintDefaults()
	}
	if help {
		flag.Usage()
		return 0, ""
	}
	return line, path
}

func tailfline(reader *bufio.Reader) {
	for {
		b, err := reader.ReadBytes('\n') // ReadBytes 读取直到输入中第一次出现分隔符，返回一个包含数据的切片，直到并包括分隔符
		if err == io.EOF {
			time.Sleep(time.Second * 1)
		}
		fmt.Print(string(b)) //
	}
}

func tailline(reader *bufio.Reader) {
	for {
		b, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(string(b))
	}
}

func tailf() {
	line, path := tailflag()
	if path == "" {
		fmt.Println("文件路径不能为空")
		return
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	// 文件不存在时报错
	if err != nil {
		return
	}
	defer file.Close()

	// 不是文件时报错
	if fileinfo, _ := file.Stat(); fileinfo.IsDir() {
		fmt.Printf("%s 为目录\n", path)
		return
	}

	// 判断
	if line != 0 {
		file.Seek(-int64(line), 2)
	}

	reader := bufio.NewReader(file)
	if line == 0 {
		tailline(reader)
	} else {
		tailfline(reader)
	}

}

func main() {
	tailf()
}
