package parser

import (
	"fmt"
	"go-crawler/engine"
	"go-crawler/model"
	"reflect"
	"regexp"
)

var regexs = map[string]*regexp.Regexp{
	"Age":        regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`),
	"Height":     regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`),
	"Weight":     regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([0-9]+)</span></td>`),
	"Marriage":   regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`),
	"Gender":     regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`),
	"Income":     regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`),
	"Education":  regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`),
	"Occupation": regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`),
	"hukou":      regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`),
	"Xinzuo":     regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`),
	"House":      regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`),
	"Car":        regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`),
}

func ParserProfile(content []byte, name string) engine.ParserResult {
	profile := model.Profile{Name: name}
	v := reflect.ValueOf(&profile).Elem()
	for k, r := range regexs {
		s := FindSubmatchContent(content, r)
		if s != "" {
			a := v.FieldByName(k)
			if a.IsValid() {
				a.Set(reflect.ValueOf(s))
			}
		} else {
			//log.Warn("未能解析的属性：%s", k)
		}
	}

	fmt.Println(profile)
	return engine.ParserResult{
		Items: []interface{}{profile},
	}
}

func FindSubmatchContent(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
