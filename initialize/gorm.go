package initialize

import (
	"fmt"
	"gin_cli/config"
	"gin_cli/model"

	gormadapter "github.com/casbin/gorm-adapter/v3"
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
	err := db.AutoMigrate(
		model.User{},
		model.Api{},
		model.Menu{},
		model.MenuParameter{},
		model.Role{},
		gormadapter.CasbinRule{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("AutoMigrate Tables Success!")
}
