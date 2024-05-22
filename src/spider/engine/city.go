package engine

import (
	"regexp"
	"strings"
)

func City(context []byte) ParseResult {
	re := regexp.MustCompile(`<a href="/(historycity/.*?)/">>>更多<<</a>`)
	res := ParseResult{}
	prefix := "https://www.tianqi24.com/"
	for _, row := range re.FindAllSubmatch(context, -1) {
		url := strings.TrimSpace(prefix + string(row[1]))
		res.Requests = append(res.Requests, Request{Url: url, FuncName: "AllMonth"})
		break
	}
	return res
}
