package models

type Repo interface {
	InsertSnippet(title, content, expires string) (int, error)
	GetUpTo10LatestSnippets() (Snippets, error)
	GetAllSnippets(limit, offset int64) (Snippets, error)
	GetSnippet(id int) (*Snippet, error)
	InsertUser(name, email, password string) error
	VerifyUser(email, password string) (int, error)
}
