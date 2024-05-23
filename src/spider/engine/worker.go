package engine

import (
	"spider/src/spider/fetcher"
)

// For simple schedulers, the distributed version is in spider_distribute below
func Worker(r Request) (ParseResult, error) {
	fetch, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return N2F[r.FuncName](fetch), nil
}
