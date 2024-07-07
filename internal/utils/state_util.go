package utils

import (
	"sync"
)

type State struct {
	mu      sync.RWMutex
	state   map[string]interface{}
	updates chan map[string]interface{}
}

func NewState(initialState map[string]interface{}) *State {
	return &State{
		state:   initialState,
		updates: make(chan map[string]interface{}),
	}
}

func (state *State) GetState() map[string]interface{} {
	state.mu.RLock()
	defer state.mu.RUnlock()
	return state.state
}

func (state *State) SetState(newState map[string]interface{}) {
	state.mu.Lock()
	defer state.mu.Unlock()
	for key, value := range newState {
		state.state[key] = value
	}
	state.updates <- state.state
}

func (state *State) Subscribe() <-chan map[string]interface{} {
	return state.updates
}
