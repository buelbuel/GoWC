package config

import (
	"os"

	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/pelletier/go-toml"
)

type StateConfig struct {
	InitialState map[string]interface{} `toml:"InitialState"`
}

func NewStateConfig() (*StateConfig, error) {
	config := &StateConfig{}
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (config *StateConfig) InitializeState() *utils.State {
	return utils.NewState(config.InitialState)
}
