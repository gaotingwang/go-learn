package parser

import (
	"regexp"
	"strconv"

	"github.com/gaotingwang/go-learn/crawler_distributed/config"

	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/model"
)

var ageRegx = regexp.MustCompile(`<span class="grayL">年龄：</span>(\d+)</td>`)
var marriageRegx = regexp.MustCompile(`<span class="grayL">婚况：</span>([^<]+)</td>`)
var heightRegx = regexp.MustCompile(`<span class="grayL">身.*高：</span>(\d+)</td>`)
var weightRegx = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var incomeRegx = regexp.MustCompile(`<span class="grayL">月.*薪：</span>([^>]+元)</td>`)
var genderRegx = regexp.MustCompile(`<span class="grayL">性别：</span>([^<]+)</td>`)
var xinzuoRegx = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^>]+)</span></td>`)
var educationRegx = regexp.MustCompile(`<span class="grayL">学.*历：</span>([^>]+)</td>`)
var occupationRegx = regexp.MustCompile(`<td><span class="label">职业： </span>([^>]+)</td>`)
var hokouRegx = regexp.MustCompile(`<span class="grayL">居住地：</span>([^>]+)</td>`)
var houseRegx = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^>]+)</span></td>`)
var CarRegx = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^>]+)</span></td>`)

var idRegx = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)

// ParseProfile 正则匹配
func parseProfile(content []byte, name string, url string, info []byte) engine.ParseResult {

	profile := model.Profile{}

	profile.Name = name
	age, err := strconv.Atoi(extractString(info, ageRegx))
	if err == nil {
		profile.Age = age
	}
	height, e := strconv.Atoi(extractString(info, heightRegx))
	if e == nil {
		profile.Height = height
	}
	weight, e := strconv.Atoi(extractString(info, weightRegx))
	if e == nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(info, genderRegx)
	profile.Income = extractString(info, incomeRegx)
	profile.Marriage = extractString(info, marriageRegx)
	profile.Education = extractString(info, educationRegx)
	profile.Occupation = extractString(info, occupationRegx)
	profile.Hokou = extractString(info, hokouRegx)
	profile.Xinzuo = extractString(info, xinzuoRegx)
	profile.House = extractString(info, houseRegx)
	profile.Car = extractString(info, CarRegx)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idRegx),
				Payload: profile,
			},
		},
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

type ProfileParser struct {
	UserUrl string
	Name    string
	Info    []byte
}

func (p ProfileParser) Parser(content []byte, _ string) engine.ParseResult {
	return parseProfile(content, p.Name, p.UserUrl, p.Info)
}

func (p ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p
}

func NewProfileParser(userUrl string, name string, info []byte) *ProfileParser {
	return &ProfileParser{UserUrl: userUrl, Name: name, Info: info}
}
