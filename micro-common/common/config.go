package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 设置前缀，不设置默认：/micro/config
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	err = config.Load(consulSource)
	return conf, err
}
