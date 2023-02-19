package server

import (
	"os"
	"strconv"
)

type IHTTPConfig interface {
	GetAddr() string
	GetReadTimeout() int
	GetWriteTimeout() int
}

type Config struct {
	addr         string
	readTimeout  int
	writeTimeout int
}

func NewHTTPConfig() IHTTPConfig {
	readTimeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"));
	if err != nil {
		panic(err)
	}

	writeTimeout, err:= strconv.Atoi(os.Getenv("APP_WRITE_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	return &Config{
		addr:         os.Getenv("APP_URL"),
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}

func (c *Config) GetAddr() string {
	return c.addr
}

func (c *Config) GetReadTimeout() int {
	return c.readTimeout
}

func (c *Config) GetWriteTimeout() int {
	return c.writeTimeout
}
