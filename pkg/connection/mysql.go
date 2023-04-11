package connection

import (
	"fmt"
	"github.com/vuho-pg/base/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(conf config.MySQL) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
