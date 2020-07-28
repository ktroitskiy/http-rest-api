package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/ktroitskiy/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, error := toml.DecodeFile(configPath, config)
	if error != nil {
		log.Fatal(error)
	}

	server := apiserver.New(config)

	if error := server.Start(); error != nil {
		log.Fatal(error)
	}
}
