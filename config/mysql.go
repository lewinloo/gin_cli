package config

import (
	"gopkg.in/ini.v1"
	"strings"
)

func LoadDatabase(file *ini.File) {
	DBType = file.Section("database").Key("DBType").String()
	DBHost := file.Section("database").Key("DBHost").String()
	DBPort := file.Section("database").Key("DBPort").String()
	DBUser := file.Section("database").Key("DBUser").String()
	DBPassword := file.Section("database").Key("DBPassword").String()
	DBName := file.Section("database").Key("DBName").String()
	DBCharset := file.Section("database").Key("DBCharset").String()

	DBDsn = strings.Join([]string{DBUser, ":", DBPassword, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=", DBCharset, "&parseTime=True&loc=Local"}, "")
}
