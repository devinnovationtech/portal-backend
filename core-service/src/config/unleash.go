package config

import (
	"github.com/spf13/viper"
)

type UnleashConfig struct {
	Url   string
	Token string
}

func LoadUnleashConfig() UnleashConfig {
	return UnleashConfig{
		Url:   viper.GetString("UNLEASH_URL"),
		Token: viper.GetString("UNLEASH_TOKEN"),
	}
}
