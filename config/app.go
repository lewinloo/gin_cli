package config

import "gopkg.in/ini.v1"

func LoadApp(file *ini.File) {
	AppMode = file.Section("app").Key("AppMode").String()
	AppPort = file.Section("app").Key("AppPort").String()
}
