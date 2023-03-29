package hystrix

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
)

type clientWrapper struct {
	client.Client
}

// 相当于一个包装类，对Client进行包装，增强了Client功能
// 熔断逻辑
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	// hystrix.Do() 同步API，第一个参数是command, 应该是与当前请求一一对应的一个名称，如入“GET-/test”。
	// 第三个参数传入一个函数，函数包含处理错误逻辑，当请求失败时应该返回error, hystrix会根据失败率执行熔断策略
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		//走熔断逻辑,每个服务可不一样
		fmt.Println(req.Service() + "." + req.Endpoint() + "的熔断逻辑")
		return e
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
