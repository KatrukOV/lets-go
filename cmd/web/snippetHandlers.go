package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrii-minchekov/lets-go/pkg/forms"
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
