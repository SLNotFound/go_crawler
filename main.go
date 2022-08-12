package main

import (
	"fmt"
	"go_crawler/fetcher"
)

//func main() {
//	url := "http://www.zhenai.com/zhenghun"
//	engine.Run(engine.Request{
//		Url:        url,
//		ParserFunc: parser.ParseCityList,
//	})
//}

func main() {
	//url := "https://album.zhenai.com/u/1213676771"
	url := "https://album.zhenai.com/u/1213676771"
	contents, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Printf("fetch url err: %v\n", err)
	}

	fmt.Println(contents)
}
