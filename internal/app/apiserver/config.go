package apiserver

// Config - структура описывающая конфиг сервера
type Config struct {
	BindAddr string `toml:"dind_addr"`
	LogLevel string `tonl:"log_level"`
}

// NewConfig - создаем новый конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}
