package engine

import (
	"spider/src/spider/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	fetch, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return N2F[r.FuncName](fetch), nil
}
