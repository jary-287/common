package consul

import (
	"fmt"

	"github.com/asim/go-micro/v3/config"
	"github.com/go-micro/plugins/v3/config/source/consul"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	source := consul.NewSource(
		consul.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		consul.WithPrefix(prefix),
	)
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	err = conf.Load(source)
	return conf, err
}
