package main

import (
	"fmt"
	"go_crawler/fetcher"
	"go_crawler/zhenai/parser"
)

func main() {
	//url := "http://www.zhenai.com/zhenghun"
	url1 := "http://www.zhenai.com/zhenghun/aba"
	contents, err := fetcher.Fetch(url1)
	if err != nil {
		fmt.Printf("fetch url err: %v\n", err)
	}
	userList1 := parser.ParseUserList1(contents)
	fmt.Println(userList1.UserUrl)
	fmt.Println(len(userList1.UserName))
}
