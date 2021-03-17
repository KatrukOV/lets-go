package main

import (
	"context"
	"github.com/andrii-minchekov/lets-go/api"
	"github.com/andrii-minchekov/lets-go/pkg/models"
)

var repo *models.Database

//ServiceGrpc allows to access Snippets from repository.
type ServiceGrpc struct{}

//ReadSnippets returns Snippets.
func (s *ServiceGrpc) ReadSnippets(ctx context.Context, in *api.PageRequest) (*api.ReadSnippetsResponse, error) {
	snippets, err := repo.GetAllSnippets(in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	apiSnippets := toApiSnippets(snippets)
	return &api.ReadSnippetsResponse{Snippets: apiSnippets}, nil
}

func toApiSnippet(snippet *models.Snippet) *api.Snippet {
	var result *api.Snippet
	result.Id = string(snippet.ID)
	result.Title = snippet.Title
	result.Content = snippet.Content
	result.Created = snippet.Created.String()
	result.Expires = snippet.Expires.String()
	return result
}

func toApiSnippets(snippets models.Snippets) []*api.Snippet {
	var result []*api.Snippet
	for _, snippet := range snippets {
		result = append(result, toApiSnippet(snippet))
	}
	return result
}
