package parser

import (
	"go-crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a></th>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/changsha/[^"]+)"`)

func ParserCity(content []byte) engine.ParserResult {
	//正则匹配出城市
	all := profileRe.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, e := range all {
		name := string(e[2])
		result.Items = append(result.Items, "User:"+name)
		result.Results = append(result.Results, engine.Request{
			Url: string(e[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParserProfile(c, name)
			},
		})
	}

	//allCity := cityUrlRe.FindAllSubmatch(content, -1)
	//for _, e := range allCity {
	//	result.Items = append(result.Items, "City:"+string(e[1]))
	//	result.Results = append(result.Results, engine.Request{
	//		Url:        string(e[1]),
	//		ParserFunc: ParserCity,
	//	})
	//}
	return result
}
