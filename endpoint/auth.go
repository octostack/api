package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/octostack/api/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	users    map[string][]byte = make(map[string][]byte)
	idxUsers int               = 0
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "cannot decode username/password struct", http.StatusBadRequest)
		return
	}

	passwordHash, found := users[u.Username]
	if !found {
		http.Error(w, "Cannot find the username", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword(passwordHash, u.Password); err != nil {
		return
	}

	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "cannot create token", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, struct {
		Token string `json:"token"`
	}{
		Token: token,
	}, http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Cannot decode request", http.StatusBadRequest)
		return
	}
	if _, found := users[u.Username]; found {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	// If I'm here-> add user and return a token
	value, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	users[u.Username] = value
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
	sendJSONResponse(w, struct {
		Token string `json:"token"`
	}{
		Token: token,
	}, http.StatusOK)
}

func Token(w http.ResponseWriter, r *http.Request) {
}

func createToken(username string) (string, error) {
	return "", nil
}
