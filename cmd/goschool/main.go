package main

import (
	"flag"
	"goschool/helpers"
	"goschool/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	// flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&configPath, "config-path", "../../configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	helpers.LogFatal(err)

	s := apiserver.New(config)
	helpers.LogFatal(s.Start())
}
