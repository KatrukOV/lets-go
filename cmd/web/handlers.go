package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrii-minchekov/lets-go/pkg/forms"
	"github.com/andrii-minchekov/lets-go/pkg/models"
)

// Home Display a "Hello from Snippetbox" message
func (app *App) Home(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	snippets, err := app.Database.GetUpTo10LatestSnippets()

	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "home.page.html", &HTMLData{
		Snippets: snippets,
	})

}

// ShowSnippet ...
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Database.GetSnippet(id)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	if snippet == nil {
		app.NotFound(w)
		return
	}

	session := app.Sessions.Load(r)
	flash, err := session.PopString(w, "flash")

	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Snippet: snippet,
		Flash:   flash,
	})
}

// NewSnippet ...
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "new.page.html", &HTMLData{
		Form: &forms.NewSnippet{},
	})
}

// CreateSnippet ...
func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	err := r.ParseForm()

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.NewSnippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: r.PostForm.Get("expires"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "new.page.html", &HTMLData{Form: form})
		return
	}

	id, err := app.Database.InsertSnippet(form.Title, form.Content, form.Expires)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	session := app.Sessions.Load(r)

	err = session.PutString(w, "flash", "Your snippet was saved!")

	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusCreated)

}

// Get all Snippets ...
func (app *App) AllSnippets(w http.ResponseWriter, r *http.Request) {

	limitStr, _ := r.URL.Query()["limit"]
	offsetStr, _ := r.URL.Query()["offset"]

	limit, _ := strconv.ParseInt(limitStr[0], 10, 64)
	offset, _ := strconv.ParseInt(offsetStr[0], 10, 64)

	snippets, err := app.Database.GetAllSnippets(limit, offset)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippets)

}

// SignupUser ...
func (app *App) SignupUser(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "signup.page.html", &HTMLData{
		Form: &forms.SignupUser{},
	})
}

// Signup ...
func (app *App) Signup(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.SignupUser{
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "signup.page.html", &HTMLData{Form: form})
		return
	}

	err = app.Database.InsertUser(form.Name, form.Email, form.Password)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	msg := "Your signup was successful. Please log in using your credentials."
	session := app.Sessions.Load(r)

	err = session.PutString(w, "flash", msg)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

// Create user
func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.UserCreate
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	err = app.Database.InsertUser(user.Name, user.Email, user.Password)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

// Login User
func (app *App) LoginUser(w http.ResponseWriter, r *http.Request) {

	var userLogin models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)

	if err != nil {
		app.ClientError(w, http.StatusUnprocessableEntity)
		return
	}
	_, err1 := app.Database.VerifyUser(userLogin.Email, userLogin.Password)

	if err1 != nil {
		app.ClientError(w, http.StatusAccepted)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// Signin ...
func (app *App) Signin(w http.ResponseWriter, r *http.Request) {

	session := app.Sessions.Load(r)

	flash, err := session.PopString(w, "flash")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "login.page.html", &HTMLData{
		Flash: flash,
		Form:  &forms.LoginUser{},
	})
}

// VerifyUser ...
func (app *App) VerifyUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.LoginUser{
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "login.page.html", &HTMLData{Form: form})
		return
	}

	currentUserID, err := app.Database.VerifyUser(form.Email, form.Password)

	if err == models.ErrInvalidCredentials {
		form.Failures["Generic"] = "Email or Password is incorrect"
		app.RenderHTML(w, r, "login.page.html", &HTMLData{Form: form})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	// Add the ID of the current user to the session
	session := app.Sessions.Load(r)
	err = session.PutInt(w, CURRENT_USER_ID, currentUserID)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/snippet/new", http.StatusFound)

}

// LogoutUser ...
func (app *App) LogoutUser(w http.ResponseWriter, r *http.Request) {

	// Remove the currentUserID from the session data.
	session := app.Sessions.Load(r)
	err := session.Remove(w, CURRENT_USER_ID)

	if err != nil {
		app.ServerError(w, err)
		return

	}

	// Redirect the user to the homepage.
	http.Redirect(w, r, "/", http.StatusResetContent)

}
