package endpoint

import (
	"net/http"
)

type AuthEndpoint struct{}

func NewAuthEndpoint() *AuthEndpoint {
	return &AuthEndpoint{}
}

func (e *AuthEndpoint) Login(w http.ResponseWriter, r *http.Request) {
}

func (e *AuthEndpoint) Register(w http.ResponseWriter, r *http.Request) {
}

func (e *AuthEndpoint) Token(w http.ResponseWriter, r *http.Request) {
}
