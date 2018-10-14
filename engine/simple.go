package engine

import (
	"errors"
	"go-crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Results...)
		for _, item := range parserResult.Items {
			log.Printf("get item %v ", item)
		}
	}
}

func Worker(r Request) (ParserResult, error) {
	if IsDuplicate(r.Url) {
		err := errors.New("is duplicate")
		return ParserResult{}, err
	}
	body, err := fetcher.Fetch(r.Url)
	log.Printf("Fetcher url:%s", r.Url)
	if err != nil {
		log.Printf("Fetcher error:%v, url:%s", err, r.Url)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
