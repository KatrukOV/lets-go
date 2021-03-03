package impl

import (
	"database/sql"
	"errors"
	"github.com/andrii-minchekov/lets-go/domain/snippet"
)

type dbSnippetRepository struct {
	db Database
}

func (r dbSnippetRepository) AddSnippet(snippet snp.Snippet) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES($1, $2, CURRENT_DATE, CURRENT_DATE + INTERVAL '100000 seconds') returning id`

	result, err := r.db.Query(stmt, snippet.Title, snippet.Content)

	if err != nil {
		return 0, err
	}

	if err := result.Next(); !err {
		return 0, errors.New("no id was generated")
	}
	var id int
	err = result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r dbSnippetRepository) LatestSnippets() (snp.Snippets, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	ORDER BY created DESC LIMIT 10`

	rows, err := r.db.Query(stmt)

	if err != nil {
		return nil, err
	}

	// This should come after we check for an error
	// or the rows object could be nil
	defer rows.Close()

	snippets := snp.Snippets{}

	for rows.Next() {

		s := &snp.Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (r dbSnippetRepository) GetSnippet(id int) (*snp.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE id = $1`

	row := r.db.QueryRow(stmt, id)

	s := &snp.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
