package config

import (
	utils "github.com/buelbuel/gowc/internal/utils"
)

func InitializeState() *utils.State {
	initialState := map[string]interface{}{
		"user":            nil,
		"isAuthenticated": false,
	}
	return utils.NewState(initialState)
}
