package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	DetailConfig detailConfig `toml:"detail"`
	TokenConfig  tokenConfig  `toml:"token"`
	MysqlConfig  mysqlConfig  `toml:"mysql"`
	RedisConfig  redisConfig  `toml:"redis"`
}

type detailConfig struct {
	PageSize      int `toml:"page_size"`
	SingleChatCap int `toml:"single_chat_cap"`
	GroupChatCap  int `toml:"group_chat_cap"`
}

type tokenConfig struct {
	SecretKey  string `toml:"secret_key"`
	EffectTime int64  `toml:"effect_time"`
}

type mysqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DNS      string
}

type redisConfig struct {
	Host string
	Port string
	DB   int
}

var Conf Config

func Init() {
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		panic(err)
	}
}
