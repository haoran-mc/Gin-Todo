package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type MySQLInfo struct {
	db   string
	user string
	pwd  string
	host string
	port int
}

func InitMySQL() (err error) {
	conf := new(MySQLInfo)
	conf.db = "Gin_Todo"
	conf.user = "root"
	conf.pwd = "haoran232"
	conf.host = "127.0.0.1"
	conf.port = 3306

	//dsn := "root:haoran232@(127.0.0.1:3306)/Gin_Todo?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.user, conf.pwd, conf.host, conf.port, conf.db)

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
