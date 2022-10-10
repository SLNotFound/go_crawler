package main

import (
	"fmt"
	"go_crawler/biaozhun/parser"
	"go_crawler/engine"
)

//var wg sync.WaitGroup
//
//func setUp(i int) {
//	defer wg.Done()
//	url := fmt.Sprintf("http://www.bzmfxz.com/biaozhun/Soft/GJBGJJYBZ/List_%d.html", i)
//	//fmt.Println(url)
//	engine.Run(engine.Request{
//		Url: url,
//		//ParserFunc: engine.NilParser,
//		ParserFunc: parser.ParseFileList,
//	})
//}

func main() {

	//url := "http://www.bzmfxz.com/biaozhun/Soft/GJBGJJYBZ/List_1.html"

	//engine.Run(engine.Request{
	//	Url: url,
	//	//ParserFunc: engine.NilParser,
	//	ParserFunc: parser.ParseFileList,
	//})
	for i := 1; i < 50; i++ {
		url := fmt.Sprintf("http://www.bzmfxz.com/biaozhun/Soft/GJBGJJYBZ/List_%d.html", i)
		fmt.Println(url)
		engine.Run(engine.Request{
			Url: url,
			//ParserFunc: engine.NilParser,
			ParserFunc: parser.ParseFileList,
		})
	}
}
