package engine

import (
	"regexp"
	"strings"
)

// AllDay output https://www.tianqi24.com/beijing/history01.html 1月
func AllDay(context []byte) (res ParseResult) {

	reRes := regexp.MustCompile(`(https://www.tianqi24.com/.+?).html`).FindSubmatch(context)
	if len(reRes) != 2 {
		return
	}
	prefix := string(reRes[1])
	re := regexp.MustCompile(`<option value='(.+?)' >(.+?)</option>`)
	for _, row := range re.FindAllSubmatch(context, -1)[5:] {
		if len(row) != 3 {
			continue
		}
		url := strings.TrimSpace(prefix + string(row[1]) + `.html`)
		month := string(row[2])
		res.Requests = append(res.Requests, Request{Url: url, FuncName: "ParserDay"})
		res.Items = append(res.Items, month)

		break
	}
	return
}
