package apiserver

import (
	"github.com/Frezyx/calory-calc-server/internal/app/store"
)

// Config - структура описывающая конфиг сервера
type Config struct {
	BindAddr string `toml:"dind_addr"`
	LogLevel string `tonl:"log_level"`
	Store    *store.Config
}

// NewConfig - создаем новый конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
