package engine

import (
	"go_crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	// 声明一个Request结构体切片
	var requests []Request

	// 循坏取出seeds里的项，加到requests里
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 循环条件为只要requests有值就执行
	for len(requests) > 0 {

		// 取出requests的第一个值
		r := requests[0]
		requests = requests[1:]

		// fetch url,拿到页面内容
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
			continue
		}
		// 对页面数据body进行parser得到 parseResult
		parseResult := r.ParserFunc(body)
		// 将parseResult的Requests加入到requests里
		requests = append(requests, parseResult.Requests...)
	}
}
