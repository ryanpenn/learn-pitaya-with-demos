package config

type ChatConfig struct {
	ServerConfig
	DbName    string
	DbAddr    string
	DbTimeout int64
}
