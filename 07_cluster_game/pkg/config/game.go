package config

type GameConfig struct {
	ServerConfig
	ServerID  int
	DbName    string
	DbAddr    string
	DbTimeout int64
}
