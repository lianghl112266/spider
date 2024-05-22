package engine

import (
	"fmt"
	"spider/src/spider/fetcher"
	"testing"
)

func TestCity(t *testing.T) {
	if context, err := fetcher.Fetch(`https://www.tianqi24.com/historycity/`); err == nil {
		City(context)
	} else {
		fmt.Println("fetch err", err)
	}
}

func TestAllMonth(t *testing.T) {
	if context, err := fetcher.Fetch(`https://www.tianqi24.com/historycity/prov_beijing`); err == nil {
		AllMonth(context)
	} else {
		fmt.Println("fetch err", err)
	}
}

func TestAllDay(t *testing.T) {
	if context, err := fetcher.Fetch(`https://www.tianqi24.com/huerle/history.html`); err == nil {
		AllDay(context)
	} else {
		fmt.Println("fetch err", err)
	}
}

func TestParserDay(t *testing.T) {
	if context, err := fetcher.Fetch(`https://www.tianqi24.com/nankaiqu/history01.html`); err == nil {
		ParserDay(context)
	} else {
		fmt.Println("fetch err", err)
	}
}
