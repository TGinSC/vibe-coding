package config

import (
	"contribution/tool"
	"encoding/hex"
)

type Config struct {
	HttpPort          string
	DB_FILE           string
	SALT_FILE         string
	SaltManager       *tool.SaltManager
	DefaultHashConfig *tool.SimpleHash
}

var Config__ *Config = ConfigInit()

func ConfigInit() *Config {
	// 初始化配置

	config := &Config{
		HttpPort:          ":1411",
		DB_FILE:           "data.db",
		SALT_FILE:         "salt.json",
		SaltManager:       tool.NewSaltManager("salt.json"),
		DefaultHashConfig: tool.NewDefaultHash(),
	}
	saltBytes, e := config.DefaultHashConfig.GenerateSalt()
	if e != nil {
		panic(e)
	}
	salt := hex.EncodeToString(saltBytes)
	e = config.SaltManager.WriteSalt(salt)
	if e != nil {
		panic(e)
	}
	return config
}
