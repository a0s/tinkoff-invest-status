package config

import (
	"flag"
	"fmt"
	"github.com/octago/sflags/gen/gflag"
	"log"
	"os"
)

const version = "0.0.1"

type Logger interface {
	Fatalln(v ...interface{})
}

type Config struct {
	Token   string `flag:"token" desc:"tinkoff REST token"`
	Version bool   `flag:"version" desc:"show version"`

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

	if config.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	if config.Token == "" {
		config.Logger.Fatalln("empty --token")
	}

	return config
}
