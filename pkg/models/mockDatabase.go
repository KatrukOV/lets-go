package models

// MockDatabase ...
type MockDatabase struct{}

// InsertSnippet ...
func (db *MockDatabase) InsertSnippet(title, content, expires string) (int, error) {
	return 1, nil
}

// GetUpTo10LatestSnippets ...
func (db *MockDatabase) GetUpTo10LatestSnippets() (Snippets, error) {
	return Snippets{&Snippet{ID: 1, Title: "Title"}}, nil
}

// GetAllSnippets ...
func (db *MockDatabase) GetAllSnippets(limit, offset int64) (Snippets, error) {
	return Snippets{&Snippet{ID: 1, Title: "Title"}}, nil
}

// GetSnippet ...
func (db *MockDatabase) GetSnippet(id int) (*Snippet, error) {
	return &Snippet{ID: 1, Title: "Title"}, nil
}

// InsertUser ...
func (db *MockDatabase) InsertUser(name, email, password string) error {
	return nil
}

// VerifyUser ...
func (db *MockDatabase) VerifyUser(email, password string) (int, error) {
	return 1, nil
}
