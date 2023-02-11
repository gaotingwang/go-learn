package parser

import (
	"fmt"
	"regexp"

	"github.com/gaotingwang/go-learn/crawler_distributed/config"

	"github.com/gaotingwang/go-learn/crawler/engine"
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
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	fmt.Println("match city count is", len(matchAll))
	return result
}
