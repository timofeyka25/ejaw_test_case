package config

import (
	"os"
	"sync"
)

type Config struct {
	ServerURL            string
	DefaultAdminUsername string
	DefaultAdminPassword string
	JwtSecret            string
}

func Get() *Config {
	var cfg Config
	var once sync.Once
	once.Do(func() {
		cfg = Config{
			ServerURL:            os.Getenv("SERVER_URL"),
			DefaultAdminUsername: os.Getenv("DEFAULT_ADMIN_USERNAME"),
			DefaultAdminPassword: os.Getenv("DEFAULT_ADMIN_PASSWORD"),
			JwtSecret:            os.Getenv("JWT_SECRET"),
		}
	})
	return &cfg
}
