package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

const cityRegx = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// ParseCity 正则匹配
func ParseCity(content []byte, replaceUrl string) engine.ParseResult {
	compile := regexp.MustCompile(cityRegx)
	matchAll := compile.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matchAll {
		fmt.Printf("User : %s, URL: %s \n", m[2], m[1])
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			//Url: string(m[1]),
			Url: replaceUrl,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				// 注意闭包问题，不能直接使用m[2]，否则最终内部函数第二个参数都是一样的
				return ParseProfile(bytes, name)
			},
		})
	}
	fmt.Println("match user count is", len(matchAll))
	return result
}
