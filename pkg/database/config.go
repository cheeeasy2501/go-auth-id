package database

import (
	"os"
)

type IConfig interface {
	getHost() string
	getUser() string
	getPassword() string
	getDBName() string
	getSSLMode() string
	getTimezone() string
	getPort() string
}

type Config struct {
	host, user, password, dbname, sslmode, timeZone, port string
}

func NewPostgresConfig() *Config {
	return &Config{
		host:     os.Getenv("POSTGRES_HOST"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
		sslmode:  os.Getenv("POSTGRES_SSLMODE"),
		timeZone: os.Getenv("POSTGRES_TIMEZONE"),
		port:     os.Getenv("POSTGRES_PORT"),
	}
}

func (c *Config) getHost() string {
	return c.host
}

func (c *Config) getUser() string {
	return c.user
}

func (c *Config) getPassword() string {
	return c.password
}

func (c *Config) getDBName() string {
	return c.dbname
}

func (c *Config) getSSLMode() string {
	return c.sslmode
}

func (c *Config) getTimezone() string {
	return c.timeZone
}

func (c *Config) getPort() string {
	return c.port
}
