package main

import (
	"encoding/json"
	"github.com/andrii-minchekov/lets-go/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test handler Of get all Snippets ...
func Test_AllSnippets(t *testing.T) {
	t.Parallel()

	//given
	app := App{Repo: &models.MockDatabase{}}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AllSnippets)

	//when
	request, err := http.NewRequest("GET", "/snippets?offset=0&limit=2", nil)
	if err != nil {
		t.Errorf("error %v", err)
	}
	handler.ServeHTTP(response, request)

	//then
	if response.Code != http.StatusOK {
		t.Fatalf("Wrong code %v", response.Code)
	}

	var snippets models.Snippets
	err = json.Unmarshal(response.Body.Bytes(), &snippets)
	if err != nil {
		return
	}
	if len(snippets) != 1 {
		t.Errorf("Size of list of snippets != 1")
	}
	if snippets[0].ID != 1 {
		t.Errorf("Id of snippet != 1")
	}

}
