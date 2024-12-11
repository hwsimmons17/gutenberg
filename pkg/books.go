package pkg

import (
	"context"
)

type BookReader interface {
	FetchBook(id int) (Book, error)
	FetchBookText(id int) (string, error)
}

type BookRepository interface {
	SaveBook(ctx context.Context, book Book) error
	SaveBookMetadata(ctx context.Context, metadata BookMetadata) error
	SaveBookNotes(ctx context.Context, notes []BookNote) error
	SaveBookSubjects(ctx context.Context, subjects []BookSubject) error
}

type Book struct {
	ID       int           `json:"id"`
	Title    string        `json:"title"`
	Author   string        `json:"author"`
	Metadata *BookMetadata `json:"metadata,omitempty"`
}

type BookMetadata struct {
	ID                  int
	BookID              int           `json:"book_id"`
	Language            *string       `json:"language"`
	Summary             *string       `json:"summary"`
	Category            *string       `json:"category"`
	ReleaseDate         *string       `json:"release_date"`
	MostRecentlyUpdated *string       `json:"most_recently_updated"`
	CopyrightStatus     *string       `json:"copyright_status"`
	Downloads           *string       `json:"downloads"`
	Notes               []BookNote    `json:"notes"`
	Subjects            []BookSubject `json:"subjects"`
}

type BookNote struct {
	ID     int
	BookID int    `json:"book_id"`
	Note   string `json:"note"`
}

type BookSubject struct {
	ID      int
	BookID  int    `json:"book_id"`
	Subject string `json:"subject"`
}
