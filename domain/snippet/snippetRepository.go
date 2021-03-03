package snp

type Repository interface {
	AddSnippet(snippet Snippet) (int, error)
	LatestSnippets() (Snippets, error)
	GetSnippet(id int) (*Snippet, error)
}
