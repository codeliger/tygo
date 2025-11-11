package bookapp

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	BookID uuid.UUID `json:"book_id"` // A line comment about BookID that should be kept.
	Title  string    `json:"title"`
	// ISBN identifier of the book, null if not known.
	ISBN ISBN `json:"isbn"`

	Genre    string    `json:"genre"    tstype:"'novel' | 'crime' | 'fantasy'"`
	Chapters []Chapter `json:"chapters"`

	PublishedAt *time.Time `json:"published_at"`
}
