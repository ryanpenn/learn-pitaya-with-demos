package config

type ServerConfig struct {
	ServerType string `yaml:"server-type"`
	IsFrontend bool   `yaml:"is-frontend"`
}
