package config

import (
    "flag"
    "github.com/spf13/viper"
)

type Config struct {
	ListenOn string `yaml:"listen_on"`
}

var configFile = flag.String("f", "etc/config.yaml", "the config file")

//NewConfig parses the config file and returns a Config struct
func NewConfig() (Config, error) {
    flag.Parse()
    viper.SetConfigFile(*configFile)
    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}