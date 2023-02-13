package client

import (
	"github.com/gaotingwang/go-learn/jsonrpc/rpcsupport"
	"github.com/gaotingwang/go-learn/jsonrpc/server"
)

func CreateProcessor(host string) (server.User, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return server.User{}, err
	}

	var result server.User
	// jsonrpc 调用返回结果，若 result.Favorite 为interface{} 会丢失原有对象类型，变为map[string]interface{}
	err = client.Call("Service.Process", "zhangsan", &result)
	if err != nil {
		return server.User{}, err
	}

	return result, nil
}
