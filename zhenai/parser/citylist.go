package parser

import (
	"go-crawler/engine"
	"regexp"
)

const re = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(content []byte) engine.ParserResult {
	//正则匹配出城市
	re := regexp.MustCompile(re)
	all := re.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	limit := 10
	for _, e := range all {
		limit--
		if limit > 5 {
			continue
		}
		if limit <= 0 {
			break
		}
		result.Items = append(result.Items, "City:"+string(e[2]))
		result.Results = append(result.Results, engine.Request{
			Url:        string(e[1]),
			ParserFunc: ParserCity,
		})
	}
	return result
}
