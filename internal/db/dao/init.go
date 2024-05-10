package dao

import (
	"github.com/MmxLearning/OpcuaServer/internal/global"
	"github.com/MmxLearning/OpcuaServer/pkg/drivers/mysql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	conf := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	var err error
	DB, err = mysql.New(&global.Config.Mysql, conf)
	if err != nil {
		log.Fatalln("连接 Mysql 失败:", err)
	}

	if err = DB.AutoMigrate(
		&User{},
	); err != nil {
		log.Fatalln("AutoMigration failed:", err)
	}
}
