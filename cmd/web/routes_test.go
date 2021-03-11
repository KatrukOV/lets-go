package main

import (
	"github.com/andrii-minchekov/lets-go/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Get all Snippets ...
func AllSnippets(t *testing.T) {
	t.Parallel()

	app := App{Database: &models.MockDatabase{}}
	hendler := http.HandlerFunc(app.AllSnippets)

	request, err := http.NewRequest("GET", "/snippets?offset=0&limit=2", nil)
	if err != nil {
		t.Errorf("error %v", err)
	}

	w := httptest.NewRecorder()

	hendler.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Fatalf("Wrong code %v", w.Code)
	}
}
