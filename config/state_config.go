package config

import (
	"os"

	utils "github.com/buelbuel/gowc/utils"
	"github.com/pelletier/go-toml"
)

// StateConfig represents the configuration for the application state.
// It includes an [InitialState] map with string keys and arbitrary values.
type StateConfig struct {
	InitialState map[string]interface{} `toml:"InitialState"`
}

// NewStateConfig creates a new instance of [StateConfig] by reading from the config.toml file.
// It returns a pointer to a [StateConfig] instance and an error if one occurs.
// It reads the config.toml file from the root directory and unmarshals the TOML data into a [StateConfig] instance.
// If the file is not found or cannot be read, it returns an error.
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

// InitializeState initializes the application state with the initial state configuration.
// It returns a pointer to a [utils.State] instance.
func (config *StateConfig) InitializeState() *utils.State {
	return utils.NewState(config.InitialState)
}
