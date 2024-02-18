package config

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	Name      string
	Container string
}

type ServerConfig struct {
	Port          string
	AuthSecretKey string
}

func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:      "localhost",
			Port:      "5432",
			User:      "postgres",
			Password:  "postgres",
			Name:      "testyfy",
			Container: "pgcontainer",
		},
		Server: ServerConfig{
			Port:          "8080",
			AuthSecretKey: "secret",
		},
	}
}
