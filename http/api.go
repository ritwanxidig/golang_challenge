package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users []User

func (a *api) getIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) postUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{ID: len(users) + 1, Name: payload.Name}
	users = append(users, u)

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	var id = len(users)
	var name = u.Name

	// check if id is already
	for _, user := range users {
		if user.ID == id {
			return errors.New("ID already exists")
		}
	}
	// check if name is provided
	if name == "" {
		return errors.New("Name is required")
	}

	users = append(users, u)
	return nil
}
