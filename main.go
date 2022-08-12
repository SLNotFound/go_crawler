package main

import (
	"go_crawler/engine"
	"go_crawler/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}

//func main() {
//	start := time.Now()
//	url := "http://www.zhenai.com/zhenghun"
//	//url1 := "http://www.zhenai.com/zhenghun/aba"
//	contents, err := fetcher.Fetch(url)
//	if err != nil {
//		fmt.Printf("fetch url err: %v\n", err)
//	}
//
//	cityList := parser.ParseCityList1(contents)
//	for _, cityUrl := range cityList.CityUrl {
//		fetch, err := fetcher.Fetch(cityUrl)
//		if err != nil {
//			fmt.Printf("fetch url err: %v\n", err)
//		}
//		userList1 := parser.ParseUserList1(fetch)
//		fmt.Println(len(userList1.UserUrl))
//	}
//	//fmt.Println(cityList.CityUrl)
//	//userList1 := parser.ParseUserList1(contents)
//	end := time.Since(start)
//	fmt.Println(end)
//}
