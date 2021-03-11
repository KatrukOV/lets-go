package main

import (
	"github.com/andrii-minchekov/lets-go/pkg/models"
	"testing"
)

// Test handler Of get all Snippets ...
func Test_AllSnippets(t *testing.T) {
	t.Parallel()

	app := App{Database: &models.MockDatabase{}}
	snippets, err := app.Database.GetAllSnippets(2, 0)

	if err != nil {
		t.Errorf("error %v", err)
	}
	snippet := snippets[0]
	if snippet.ID != 1 {
		t.Errorf("not equal to 1 : %v", snippet.ID)
	}

}
