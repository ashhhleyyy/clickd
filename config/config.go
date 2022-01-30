package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	AllowedHosts []string
	DbDsn        string
	FollowGPC    bool
	Host         string
}

var Configuration Config

func init() {
	config, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	_, err = toml.Decode(string(config), &Configuration)
	if err != nil {
		panic(err)
	}
}
