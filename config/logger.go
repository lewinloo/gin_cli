package config

import "gopkg.in/ini.v1"

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	Compress   bool   `json:"compress"`
}

func (config *LogConfig) Load(file *ini.File) {
	config.Level = file.Section("logger").Key("Level").String()
	config.Filename = file.Section("logger").Key("Filename").String()
	config.MaxAge, _ = file.Section("logger").Key("MaxAge").Int()
	config.MaxSize, _ = file.Section("logger").Key("MaxSize").Int()
	config.Compress, _ = file.Section("logger").Key("Compress").Bool()
	config.MaxBackups, _ = file.Section("logger").Key("MaxBackups").Int()
}
