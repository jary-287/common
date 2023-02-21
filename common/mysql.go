package common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	PassWord string `json:"password"`
	User     string `json:"user"`
	Database string `json:"database"`
}

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlinfo := &MysqlConfig{}
	config.Get(path...).Scan(mysqlinfo)
	return mysqlinfo
}
