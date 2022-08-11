package main

import (
	"fmt"
	"go_crawler/fetcher"
	"go_crawler/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	contents, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Printf("fetch url err: %v\n", err)
	}
	parser.ParseCityList1(contents)
}
