package parser

import (
	"go_crawler/engine"
	"regexp"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListReg)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
