package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gaotingwang/go-learn/micro-common/common"
	hystrix2 "github.com/gaotingwang/go-learn/micro-v3/plugin/hystrix"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
)

func main() {
	// 1. 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"localhost:8500",
		}
	})

	// 2. 配置中心
	conf, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		fmt.Println(err)
	}

	// 3. 加载配置
	mysqlConfig := common.GetMysqlFromConsul(conf, "mysql")
	// 初始化数据库
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Pwd+"@/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接mysql 成功")
	defer db.Close()
	db.SingularTable(true)

	// 4. 添加链路追踪
	t, io, err := common.NewTracer("base", "localhost:6831")
	if err != nil {
		fmt.Println(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 5. 添加熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动监听程序
	go func() {
		// http://宿主机ip:9092/turbine/turbine.stream
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9092"), hystrixStreamHandler)
		fmt.Println("333")
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 6. 添加日志中心

	// 7. 添加监控

	// 创建服务
	service := micro.NewService(
		micro.Name("base"),
		micro.Version("latest"),
		// 添加注册中心
		micro.Registry(consulRegistry),
		micro.Config(conf),
		// 添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 客户端添加熔断
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		// 服务端添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
	)

	// 初始化服务
	service.Init()

	// 启动服务
	if err := service.Run(); err != nil {
		//输出启动失败信息
		log.Fatal(err)
	}
}
