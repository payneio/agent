package config

import (
	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Version string
	Process AgentProcess
}

type AgentProcess struct {
	Start   string
	Stop    string
	Restart string
	PidFile string
}

var Config Configuration
var ConfigFilename string

func Load() error {
	if _, err := toml.DecodeFile(ConfigFilename, &Config); err != nil {
		return err
	}
	return nil
}
