package parser

import (
	"fmt"
	"go_crawler/engine"
	"regexp"
)

const fileUrlReg = `<a href="(http://down.bzmfxz.com/[0-9]+/[0-9a-zA-Z]+/[0-9a-zA-Z]+.rar)" style="color:#0066FF">点击下载</a>`

func ParseFile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(fileUrlReg)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[0]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	//filePath := "F:\\biaozhun\\link.txt"
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	fmt.Printf("open file err: %v\n", err)
	//}
	//defer file.Close()
	//
	//writer := bufio.NewWriter(file)
	//
	for _, v := range result.Requests {

		//fmt.Printf("%s write to file successed\n", v.Url)
		fmt.Println(v.Url)
	}
	//
	//writer.Flush()
	return result
}
