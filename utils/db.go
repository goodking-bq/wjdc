package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type MysqlConfig struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	Options  string
}

func (mc MysqlConfig) Dsn() string {
	//root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", mc.User, mc.Password, mc.Host, mc.Port, mc.DB, mc.Options)
}

var DB *gorm.DB

func InitDB(cfg MysqlConfig) {
	db, err := gorm.Open(mysql.Open(cfg.Dsn()), &gorm.Config{PrepareStmt: false,
		DisableForeignKeyConstraintWhenMigrating: true,

		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		}})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

}
