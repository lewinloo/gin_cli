package config

import "gopkg.in/ini.v1"

type CasbinConfig struct {
	ModelPath string
}

func (conf *CasbinConfig) Load(file *ini.File) {
	conf.ModelPath = file.Section("casbin").Key("ModelPath").String()
}
