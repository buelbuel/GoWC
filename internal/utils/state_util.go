package utils

import (
	"sync"
)

// State represents a thread-safe state container with update notifications.
// It uses a read-write mutex for concurrent access and a channel for updates.
type State struct {
	mu      sync.RWMutex                // Protects concurrent access to the state
	state   map[string]interface{}      // Holds the actual state data
	updates chan map[string]interface{} // Channel for broadcasting state updates
}

// NewState creates and initializes a new State instance.
// It takes an initial state and sets up the update channel.
//
// Parameters:
//   - initialState: A map representing the initial state values
//
// Returns:
//   - A pointer to the newly created State instance
func NewState(initialState map[string]interface{}) *State {
	return &State{
		state:   initialState,
		updates: make(chan map[string]interface{}),
	}
}

// GetState retrieves the current state.
// It uses a read lock to ensure thread-safe access to the state.
//
// Returns:
//   - A copy of the current state map
func (state *State) GetState() map[string]interface{} {
	state.mu.RLock()
	defer state.mu.RUnlock()
	return state.state
}

// SetState updates the state with new values.
// It acquires a write lock, updates the state, and broadcasts the change.
//
// Parameters:
//   - newState: A map containing the new state values to be merged
//
// Note: This method sends the entire state to the updates channel,
// which might be inefficient for large states or frequent updates.
func (state *State) SetState(newState map[string]interface{}) {
	state.mu.Lock()
	defer state.mu.Unlock()
	for key, value := range newState {
		state.state[key] = value
	}
	state.updates <- state.state
}

// Subscribe returns a read-only channel that receives state updates.
// Clients can use this to be notified of any changes to the state.
//
// Returns:
//   - A receive-only channel of map[string]interface{}
//
// Note: Proper handling of this channel is crucial to prevent goroutine leaks.
// Clients should ensure they continue reading from this channel or close it when no longer needed.
func (state *State) Subscribe() <-chan map[string]interface{} {
	return state.updates
}
