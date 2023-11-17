package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Load[T any](path, name, typ string) (*T, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(typ)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("load config err: %v", err)
	}

	var c T
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unmarshal config err: %v", err)
	}
	return &c, nil
}
