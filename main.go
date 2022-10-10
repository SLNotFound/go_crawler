package main

import (
	"bufio"
	"fmt"
	"go_crawler/fetcher"
	"io"
	"os"
	"regexp"
	"strings"
)

const fileUrlReg = `<a href="(http://down.bzmfxz.com/[0-9]+/[0-9a-zA-Z]+/[0-9a-zA-Z]+.[0-9]+.rar)" style="color:#0066FF">点击下载</a>`

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读取完毕！")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}

		fmt.Print(line)
	}
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				if len(line) != 0 {
					fmt.Println(line)
				}
				fmt.Println("文件读取完毕！")
				break
			}
			ch1 <- line
			if err != nil {
				fmt.Println("read file failed, err: ", err)
				return
			}

		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i
		}
		close(ch2)
	}()

	for i := range ch2 {
		contents, err := fetcher.Fetch(strings.Trim(i, "\r\n"))
		if err != nil {
			fmt.Printf("fetch url failed, err: %v\n", err)
			continue
		}
		re := regexp.MustCompile(fileUrlReg)
		matches := re.FindAllSubmatch(contents, -1)
		for _, m := range matches {
			fmt.Println(string(m[1]))
		}
	}
}
