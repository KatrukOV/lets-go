package uc

import (
	"errors"
	snp "github.com/andrii-minchekov/lets-go/domain/snippet"
	"log"
	"reflect"
	"testing"
)

type mockRepo struct {
	expectedErr  error
	expSnippetId int
}

func (r mockRepo) LatestSnippets() (snp.Snippets, error) {
	panic("implement me")
}

func (r mockRepo) GetSnippet(id int) (*snp.Snippet, error) {
	panic("implement me")
}

func (r mockRepo) AddSnippet(snippet snp.Snippet) (int, error) {
	log.Printf("Stub AddSnippet is called")
	if r.expectedErr != nil {
		return 0, r.expectedErr
	}
	return r.expSnippetId, nil
}

func TestSnippetUseCase_CreateSnippet(t *testing.T) {
	type fields struct {
		Repo snp.Repository
	}
	type args struct {
		snippet snp.Snippet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			"should create snippet successfully when snippet input is valid",
			fields{Repo: mockRepo{expSnippetId: 1}},
			args{snp.Snippet{}},
			1,
			false,
		},
		{
			"should return expectedErr when repo fails to add snippet",
			fields{Repo: mockRepo{expectedErr: errors.New("couldn't add snippet")}},
			args{snp.Snippet{}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewSnippetUseCase(tt.fields.Repo)
			got, err := useCase.CreateSnippet(tt.args.snippet)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSnippet() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSnippet() got = %v, wantValue %v", got, tt.want)
			}
		})
	}
}

func TestSnippetUseCase_GetSnippet(t *testing.T) {
	type fields struct {
		Repo snp.Repository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *snp.Snippet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewSnippetUseCase(tt.fields.Repo)
			got, err := useCase.GetSnippet(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSnippet() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSnippet() got = %v, wantValue %v", got, tt.want)
			}
		})
	}
}

func TestSnippetUseCase_LatestSnippets(t *testing.T) {
	type fields struct {
		Repo snp.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    snp.Snippets
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewSnippetUseCase(tt.fields.Repo)
			got, err := useCase.LatestSnippets()
			if (err != nil) != tt.wantErr {
				t.Errorf("LatestSnippets() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LatestSnippets() got = %v, wantValue %v", got, tt.want)
			}
		})
	}
}
