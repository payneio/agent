package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Process
}

type Process struct {
	Start     string
	Stop      string
	Restart   string
	Something int
}

var Config = new(Config)
var ConfigFilename string

func Load() error {
	if _, err := toml.DecodeFile(ConfigFilename, &Config); err != nil {
		return err
	}
	return nil
}
