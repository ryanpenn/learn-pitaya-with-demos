package config

type LoginConfig struct {
	ServerConfig
	HttpPort    int
	Mode        string
	HttpTimeout int64
	ContextPath string
}
