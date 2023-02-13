package server

import (
	"github.com/gaotingwang/go-learn/jsonrpc/rpcsupport"
	"log"
)

type Service struct {
}

// rpc服务接口定义时入参需遵循jsonrpc的定义规范：
//  1. 方法必须具备两个参数和一个error类型的返回值
//     1.1 第一个参数为客户端调用RPC时交给服务端的数据，可以是指针也可以是实体
//     1.2 第二个参数为服务端返回给客户端的数据
//  2. 返回值需以指针方式传
func (Service) Process(name string, result *User) error {
	some := NewSome("favorite", "play")
	user := NewUser(name, 18, []byte("hello world"), some)

	*result = user
	return nil
}

func Rpc(host string, service Service) {
	log.Fatal(rpcsupport.ServerRpc(host, service))
}
