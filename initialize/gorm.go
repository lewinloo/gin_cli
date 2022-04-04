package initialize

import (
	"fmt"
	"gin_cli/config"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch config.DBType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func GenerateTables(db *gorm.DB) {
	err := db.AutoMigrate()

	if err != nil {
		panic(err)
	}

	fmt.Println("AutoMigrate Tables Success!")
}
