package server

import (
	"os"
	"strconv"
)

type IHTTPConfig interface {
	GetHost() string
	GetPort() string
	GetAddr() string
	GetReadTimeout() int
	GetWriteTimeout() int
}

type Config struct {
	host         string
	port         string
	readTimeout  int
	writeTimeout int
}

func NewHTTPConfig() IHTTPConfig {
	readTimeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	writeTimeout, err := strconv.Atoi(os.Getenv("APP_WRITE_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	return &Config{
		host:         os.Getenv("APP_HOST"),
		port:         os.Getenv("APP_PORT"),
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}

func (c *Config) GetHost() string {
	return c.host
}

func (c *Config) GetPort() string {
	return c.port
}

func (c *Config) GetAddr() string {
	return c.host + ":" + c.port
}

func (c *Config) GetReadTimeout() int {
	return c.readTimeout
}

func (c *Config) GetWriteTimeout() int {
	return c.writeTimeout
}
