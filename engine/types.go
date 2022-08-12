package engine

// Request结构体里存url和url的parser函数
// parser函数传入[]byte 返回ParseResult
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
