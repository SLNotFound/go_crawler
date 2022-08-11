package parser

import (
	"go_crawler/model"
	"regexp"
)

const testReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList1(contents []byte) model.City {
	re := regexp.MustCompile(testReg)
	matches := re.FindAllSubmatch(contents, -1)

	result := model.City{}
	limit := 2
	for _, match := range matches {
		result.CityName = append(result.CityName, string(match[2]))
		result.CityUrl = append(result.CityUrl, string(match[1])+"\n")
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
