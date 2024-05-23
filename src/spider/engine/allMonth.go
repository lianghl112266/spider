package engine

import (
	"regexp"
	"strings"
)

// AllMonth https://www.tianqi24.com//beijing/history.html 北京
func AllMonth(context []byte) (res ParseResult) {
	re := regexp.MustCompile(`<a href="(/(.+?)/history.html)" title.+?>(.+?)</a>`)
	prefix := "https://www.tianqi24.com"
	for i, row := range re.FindAllSubmatch(context, -1) {
		if i == 0 {
			continue
		}
		url := strings.TrimSpace(prefix + string(row[1]))
		city := string(row[2])
		res.Items = append(res.Items, city)
		res.Requests = append(res.Requests, Request{Url: url, FuncName: "AllDay"})

		break
	}
	return
}
