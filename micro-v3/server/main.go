package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gaotingwang/go-learn/micro-v3/common"
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
		common.Error(err)
	}

	// 3. 加载配置
	mysqlConfig := common.GetMysqlFromConsul(conf, "mysql")
	// 初始化数据库
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Pwd+"@/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		common.Error(err)
	}
	common.Debug("连接mysql 成功")
	defer db.Close()
	db.SingularTable(true)

	// 4. 添加链路追踪
	t, io, err := common.NewTracer("base", "localhost:6831")
	if err != nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 5. 添加熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动监听程序
	go func() {
		// http://宿主机ip:9092/turbine/turbine.stream
		err = http.ListenAndServe(net.JoinHostPort("10.64.86.100", "9092"), hystrixStreamHandler)
		if err != nil {
			common.Error(err)
		}
	}()

	// 6. 暴露端口给prometheus采集信息
	common.PrometheusBoot(9093)

	// 7. 添加日志中心（日志会写入micro.log文件中，启动filebeat上传文件到logstash中）
	common.Debug("Debug 日志")
	common.Info("Info 日志")
	common.Error("Error 日志")

	// 创建服务
	service := micro.NewService(
		micro.Name("base"),
		micro.Version("latest"),
		// 设置服务地址和需要暴露的端口
		micro.Address("127.0.0.1:8081"),
		// 添加注册中心
		micro.Registry(consulRegistry),
		// 添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 客户端添加熔断
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		// 服务端添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
		// 添加监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// 初始化服务
	service.Init()

	// 启动服务
	if err := service.Run(); err != nil {
		//输出启动失败信息
		common.Error(err)
	}
}
