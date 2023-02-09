package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRegx = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d])+岁</td>`)
var marriageRegx = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRegx = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var weightRegx = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var incomeRegx = regexp.MustCompile(`<td><span class="label">月收入：</span>([^>]+)</td>`)
var genderRegx = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
var xinzuoRegx = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^>]+)</span></td>`)
var educationRegx = regexp.MustCompile(`<td><span class="label">学历：</span>([^>]+)</td>`)
var occupationRegx = regexp.MustCompile(`<td><span class="label">职业： </span>([^>]+)</td>`)
var hokouRegx = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^>]+)</td>`)
var houseRegx = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^>]+)</span></td>`)
var CarRegx = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^>]+)</span></td>`)

// ParseProfile 正则匹配
func ParseProfile(content []byte, name string) engine.ParseResult {

	profile := model.Profile{}

	profile.Name = name
	age, err := strconv.Atoi(extractString(content, ageRegx))
	if err == nil {
		profile.Age = age
	}
	height, e := strconv.Atoi(extractString(content, heightRegx))
	if e == nil {
		profile.Height = height
	}
	weight, e := strconv.Atoi(extractString(content, weightRegx))
	if e == nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(content, genderRegx)
	profile.Income = extractString(content, incomeRegx)
	profile.Marriage = extractString(content, marriageRegx)
	profile.Education = extractString(content, educationRegx)
	profile.Occupation = extractString(content, occupationRegx)
	profile.Hokou = extractString(content, hokouRegx)
	profile.Xinzuo = extractString(content, xinzuoRegx)
	profile.House = extractString(content, houseRegx)
	profile.Car = extractString(content, CarRegx)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(content []byte, regx *regexp.Regexp) string {
	match := regx.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
