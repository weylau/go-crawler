package engine

import (
	"go-crawler-zhenai/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		log.Printf("Fetcher url:%s", body)
		return
		log.Printf("Fetcher url:%s", r.Url)
		if err != nil {
			log.Printf("Fetcher error:%v, url:%s", err, r.Url)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Results...)
		for _, item := range parserResult.Items {
			log.Printf("get item %s ", item)
		}
	}
}
