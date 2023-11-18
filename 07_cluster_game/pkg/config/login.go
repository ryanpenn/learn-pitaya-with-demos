package config

type LoginConfig struct {
	ServerConfig
	HttpPort    int
	Mode        string
	HttpTimeout int64
	TokenSecure string
	ContextPath string
	DbName      string
	DbAddr      string
	DbTimeout   int64
}
