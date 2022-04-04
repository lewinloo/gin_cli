package initialize

import (
	"gin_cli/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接数据库的配置
func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:               config.DBDsn,
		DefaultStringSize: 191,
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: false,
		//NamingStrategy: schema.NamingStrategy{
		//	//TablePrefix:   "t_",  // 表面前缀，如：t_users
		//	SingularTable: false, // 去掉复数，使用单数表名，如：t_user => t_users
		//},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
