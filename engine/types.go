package engine

// Request object, containing the URL, and the name of the parser used
type Request struct {
	Url      string
	FuncName string
}

// The object parsed by the parser
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// Empty parser, used when there is a URL but no parsing is needed
func NilParser(_ []byte) ParseResult { return ParseResult{} }

// Mapping of parser function name to function entity
// Q: Why does Request not use the function entity directly,
// but uses the name and then maps it?
// A: Because rpc cannot pass function types
var N2F = map[string]func([]byte) ParseResult{
	"City":      City,
	"AllDay":    AllDay,
	"AllMonth":  AllMonth,
	"ParserDay": ParserDay,
	"Nil":       NilParser,
}
