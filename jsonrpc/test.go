package main

import (
	"fmt"
	"github.com/gaotingwang/go-learn/jsonrpc/client"
	"github.com/gaotingwang/go-learn/jsonrpc/server"
	"time"
)

func main() {
	// 服务端
	const host = ":1234"
	go server.Rpc(host, server.Service{})
	time.Sleep(time.Second) //保证服务起来

	// 客户端
	user, err := client.CreateProcessor()
	if err != nil {
		panic(err)
	}

	// rpc 接口调用返回结果，对json反序列化时，若对象属性中有interface{}类型，会映射为map[string]interface{}
	if m, ok := user.Favorite.(map[string]interface{}); ok {
		fmt.Printf("%s\n", m["Name"].(string))
		if s, ok := m["Type"].(string); ok {
			fmt.Printf("%s\n", s)
		}
	}
	fmt.Printf("%+v", user)
}
