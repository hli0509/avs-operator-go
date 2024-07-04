package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Provider                 string
	PrivateKey               string
	ContractAddress          string
	DelegationManagerAddress string
	StakeRegistryAddress     string
	AvsDirectoryAddress      string
}

func LoadConfig() Config {

	var configPath = "config.toml"

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
