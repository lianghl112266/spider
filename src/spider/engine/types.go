package engine

type Request struct {
	Url      string
	FuncName string
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(_ []byte) ParseResult { return ParseResult{} }

var N2F = map[string]func([]byte) ParseResult{
	"City":      City,
	"AllDay":    AllDay,
	"AllMonth":  AllMonth,
	"ParserDay": ParserDay,
	"Nil":       NilParser,
}
