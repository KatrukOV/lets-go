package main

import (
	"net/http"
)

//CurrentUserID
const CURRENT_USER_ID = "currentUserID"

// LoggedIn ...
func (app *App) LoggedIn(r *http.Request) (bool, error) {

	session := app.Sessions.Load(r)
	loggedIn, err := session.Exists(CURRENT_USER_ID)

	if err != nil {
		return false, err
	}

	return loggedIn, nil
}
