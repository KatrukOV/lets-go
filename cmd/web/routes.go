package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

// Routes ...
func (app *App) Routes() http.Handler {

	mux := pat.New()

	// mux := http.NewServeMux()

	mux.Get("/", NoSurf(app.Home))
	mux.Get("/snippet/new", app.RequireLogin(NoSurf(app.NewSnippetView)))
	mux.Post("/snippet/new", app.RequireLogin(NoSurf(app.CreateSnippet)))
	mux.Get("/snippet/:id", NoSurf(app.ShowSnippet))

	// Application Routes [userUseCase]
	mux.Get("/user/signup", NoSurf(app.SignupUserView))
	mux.Post("/user/signup", NoSurf(app.SignupUser))
	mux.Get("/user/login", NoSurf(app.SigninUserView))
	mux.Post("/user/login", NoSurf(app.SignInUser))
	mux.Post("/user/logout", app.RequireLogin(NoSurf(app.LogoutUser)))

	fileServer := http.FileServer(http.Dir(app.Config.StaticDir()))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return LogRequest(SecureHeaders(mux))
}
