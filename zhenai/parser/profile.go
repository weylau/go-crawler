package parser

import (
	"go-crawler-zhenai/engine"
	"regexp"
	"strconv"
)

const reProfileAge = `<td><span class="label">年龄：</span>([\d])岁</td>`

func ParserProfile(content []byte) engine.ParserResult {
	//正则匹配出城市
	re := regexp.MustCompile(reProfileAge)
	ageMatch := re.FindSubmatch(content)
	if ageMatch != nil {
		age, err := strconv.Atoi(string(ageMatch[1]))
		if err != nil {

		}
	}

	result := engine.ParserResult{}
	for _, e := range all {
		result.Items = append(result.Items, string(e[2]))
		result.Results = append(result.Results, engine.Request{
			Url:        string(e[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
