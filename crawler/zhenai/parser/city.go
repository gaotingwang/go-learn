package parser

import (
	"fmt"
	"github.com/gaotingwang/go-learn/crawler/engine"
	"regexp"
)

var profileRegx = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>(.*?)</a></div>`)

// ParseCity 正则匹配
func ParseCity(content []byte, _ string) engine.ParseResult {
	matchAll := profileRegx.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matchAll {
		fmt.Printf("User : %s, URL: %s \n", m[2], m[1])
		name := string(m[2])
		userInfo := m[3]
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: "", // 适配新版，用户信息不再请求（403需要登录）改为直接从当前页获取
			//ParserFunc: func(bytes []byte) engine.ParseResult {
			//	// 注意闭包问题，不能直接使用m[2]，否则最终内部函数第二个参数都是一样的
			//	return ParseProfile(bytes, name, url, userInfo)
			//},
			// 这里不用担心闭包问题，函数调用的参数本身就是一个拷贝
			ParserFunc: ProfileParser(string(m[1]), name, userInfo),
		})
	}
	fmt.Println("match user count is", len(matchAll))
	return result
}
