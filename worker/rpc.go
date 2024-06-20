package worker

import (
	"spider/engine"
	"spider/fetcher"
)

type WorkerService struct {
}

// Worker provided by the worker server to save data
func (_ *WorkerService) Worker(r engine.Request, res *engine.ParseResult) error {
	fetch, err := fetcher.Fetch(r.Url)
	if err != nil {
		*res = engine.ParseResult{}
		return err
	}
	*res = engine.N2F[r.FuncName](fetch)
	return nil
}
