package parser

import (
	"fmt"
	"go_crawler/engine"
	"regexp"
)

const URL = "http://www.bzmfxz.com/Common/ShowDownloadUrl.aspx?urlid=0&id="

const fileListReg = `<a href="/biaozhun/Soft/GJBGJJYBZ/[0-9]+/[0-9]+/[0-9]+/([0-9]+).html" target="_blank">([^<]+)</a>`

//const fileListReg = `<a href="(/biaozhun/Soft/GJBGJJYBZ/[0-9]+/[0-9]+/[0-9]+/[0-9]+.html)" target="_blank">([^<]+)</a>`

func ParseFileList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(fileListReg)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: fmt.Sprintf("%s%s", URL, string(m[1])),
			//ParserFunc: engine.NilParser,
			ParserFunc: ParseFile,
		})
	}

	for _, v := range result.Requests {
		fmt.Println(v.Url)
	}

	return result
}
