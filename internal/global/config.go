package global

import "github.com/MmxLearning/OpcuaServer/pkg/drivers/mysql"

var Config _Config

type _Config struct {
	Mysql mysql.Config `yaml:"mysql"`
}
