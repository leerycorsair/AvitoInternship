package config

type RepositoryConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type ServerConfig struct {
	StartPort string           `toml:"start_port"`
	RepConfig RepositoryConfig `toml:"server"`
}

func CreateServerConfig() *ServerConfig {
	return &ServerConfig{}
}
