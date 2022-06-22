package endpoint

import "net/http"

type UserEndpoint struct{}

func NewUserEndpoint() *UserEndpoint {
	return &UserEndpoint{}
}

// List users
func (e *UserEndpoint) List(w http.ResponseWriter, r *http.Request) {
}

// Get a user
func (e *UserEndpoint) Get(w http.ResponseWriter, r *http.Request) {
}

// Update a user
func (e *UserEndpoint) Update(w http.ResponseWriter, r *http.Request) {
}

// Delete a user
func (e *UserEndpoint) Delete(w http.ResponseWriter, r *http.Request) {
}

// Get contextual information for a user
func (e *UserEndpoint) GetHovercard(w http.ResponseWriter, r *http.Request) {
}
