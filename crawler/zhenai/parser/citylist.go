package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

const cityListRegx = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList 正则匹配
func ParseCityList(content []byte, _ string) engine.ParseResult {
	compile := regexp.MustCompile(cityListRegx)
	matchAll := compile.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matchAll {
		//fmt.Printf("City : %s, URL: %s \n", m[2], m[1])
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	fmt.Println("match city count is", len(matchAll))
	return result
}
