package parser

import (
	"go-crawler/engine"
	"regexp"
)

const profileRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a></th>`

func ParserCity(content []byte) engine.ParserResult {
	//正则匹配出城市
	re := regexp.MustCompile(profileRe)
	all := re.FindAllSubmatch(content, -1)
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
	return result
}
