package endpoint

import "net/http"

type OrgEndpoint struct{}

func NewOrgEndpoint() *OrgEndpoint {
	return &OrgEndpoint{}
}

func (e *OrgEndpoint) List(w http.ResponseWriter, r *http.Request) {
}

func (e *OrgEndpoint) Get(w http.ResponseWriter, r *http.Request) {
}
