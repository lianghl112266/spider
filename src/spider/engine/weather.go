package engine

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"regexp"
	"spider/src/spider/module"
	"strings"
)

func ParserDay(context []byte) ParseResult {
	res := ParseResult{}
	re := regexp.MustCompile(`<sup>(.*?)历史天气</sup>`)
	province := re.FindSubmatch(context)
	if len(province) != 2 {
		fmt.Println(province)
		fmt.Println(string(context))
		return res
	}

	root, _ := htmlquery.Parse(bytes.NewReader(context))
	for _, li := range htmlquery.Find(root, `//*[@id="main"]/section/article[3]/section/ul/li`)[1:] {
		devs := htmlquery.Find(li, `./div`)

		w := module.Weather{
			Province: string(province[1]),
			Date:     strings.TrimSpace(htmlquery.InnerText(devs[0])),
			Weather: func(s string) string {
				s = strings.Replace(s, "\n", "", -1)
				//s = strings.Replace(s, " ", "", -1)
				ss := strings.Split(s, "/")
				for i := range ss {
					ss[i] = strings.TrimSpace(ss[i])
				}
				s = strings.Join(ss, "/")
				return s
			}(strings.TrimSpace(htmlquery.InnerText(devs[1]))),
			HighTemperature: strings.TrimSpace(htmlquery.InnerText(devs[2])),
			LowTemperature:  strings.TrimSpace(htmlquery.InnerText(devs[3])),
			AQI:             strings.TrimSpace(htmlquery.InnerText(devs[4])),
			Wind:            strings.TrimSpace(htmlquery.InnerText(devs[5])),
			Rain:            strings.TrimSpace(htmlquery.InnerText(devs[6])),
		}

		fmt.Println(w)
		res.Requests = append(res.Requests, Request{Url: "", FuncName: "NilParser"})
		res.Items = append(res.Items, w)
	}
	return res
}
