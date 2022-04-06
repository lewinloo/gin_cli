package config

import "gopkg.in/ini.v1"

type JwtConfig struct {
	JwtSecret  string
	ExpireTime int
	Issuer     string
	Subject    string
}

func (conf *JwtConfig) Load(file *ini.File) {
	conf.JwtSecret = file.Section("jwt").Key("JwtSecret").String()
	conf.ExpireTime, _ = file.Section("jwt").Key("ExpireTime").Int()
	conf.Issuer = file.Section("jwt").Key("Issuer").String()
	conf.Subject = file.Section("jwt").Key("Subject").String()
}
