package config

import (
	"flag"
	"github.com/octago/sflags/gen/gflag"
	"log"
	"os"
)

type Logger interface {
	Fatalln(v ...interface{})
}

type Config struct {
	Token   string `flag:"token" desc:"tinkoff REST token"`

	Logger Logger
}

func NewConfig() *Config {
	config := &Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	err := gflag.ParseToDef(config)
	if err != nil {
		config.Logger.Fatalln("new config:", err)
	}
	flag.Parse()

	if config.Token == "" {
		config.Logger.Fatalln("empty --token")
	}

	return config
}
