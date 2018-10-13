package parser

import (
	"fmt"
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
	"strconv"
)

var (
	ageRe        = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([0-9]+)</span></td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	nameRe       = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hukouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xinzuoRe     = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
)

func ParserProfile(content []byte, name string) engine.ParserResult {
	profile := model.Profile{Name: name}
	age, err := strconv.Atoi(FindSubmatchContent(content, ageRe))
	fmt.Println(age)
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(FindSubmatchContent(content, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(FindSubmatchContent(content, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Marriage = FindSubmatchContent(content, marriageRe)
	profile.Name = FindSubmatchContent(content, nameRe)
	profile.Gender = FindSubmatchContent(content, genderRe)
	profile.Income = FindSubmatchContent(content, incomeRe)
	profile.Education = FindSubmatchContent(content, educationRe)
	profile.Occupation = FindSubmatchContent(content, occupationRe)
	profile.Hukou = FindSubmatchContent(content, hukouRe)
	profile.Xinzuo = FindSubmatchContent(content, xinzuoRe)
	profile.House = FindSubmatchContent(content, houseRe)
	profile.Car = FindSubmatchContent(content, carRe)
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
