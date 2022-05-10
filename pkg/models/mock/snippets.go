package mock

import (
	"time"

	"example.com/snippetbox/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Hola Mundo",
	Content: "This is a mock...",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetModel ...
type SnippetModel struct{}

// Insert ...
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

// Get ...
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecords
	}
}

// Latest ...
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
