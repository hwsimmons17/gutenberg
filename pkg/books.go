package pkg

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BookWithText struct {
	Book Book   `json:"book"`
	Text string `json:"text"`
}

type Book struct {
	bun.BaseModel `bun:"table:books"`
	ID            int           `json:"id,omitempty" bun:",pk,autoincrement"`
	CreatedAt     time.Time     `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	Title         string        `json:"title"`
	Author        string        `json:"author"`
	Metadata      *BookMetadata `json:"metadata,omitempty" bun:"rel:has-one,join:id=book_id"`
}

type BookMetadata struct {
	bun.BaseModel       `bun:"table:book_metadatas"`
	ID                  int           `bun:",pk,autoincrement"`
	CreatedAt           time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
	BookID              int           `json:"book_id"`
	Language            *string       `json:"language"`
	Summary             *string       `json:"summary"`
	Category            *string       `json:"category"`
	ReleaseDate         *string       `json:"release_date"`
	MostRecentlyUpdated *string       `json:"most_recently_updated"`
	CopyrightStatus     *string       `json:"copyright_status"`
	Downloads           *string       `json:"downloads"`
	Notes               []BookNote    `json:"notes" bun:"rel:has-many,join:book_id=book_id"`
	Subjects            []BookSubject `json:"subjects" bun:"rel:has-many,join:book_id=book_id"`
}

type BookNote struct {
	bun.BaseModel `bun:"table:book_notes"`
	ID            int       `bun:",pk,autoincrement"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	BookID        int       `json:"book_id"`
	Note          string    `json:"note"`
}

type BookSubject struct {
	bun.BaseModel `bun:"table:book_subjects"`
	ID            int       `bun:",pk,autoincrement"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	BookID        int       `json:"book_id"`
	Subject       string    `json:"subject"`
}

type UserBook struct {
	bun.BaseModel `bun:"table:user_books"`
	ID            int       `bun:",pk,autoincrement"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UserID        uuid.UUID `json:"user_id"`
	BookID        int       `json:"book_id"`
}

type BookReader interface {
	FetchBook(id int) (Book, error)
	FetchBookText(id int) (string, error)
}

type BookRepository interface {
	ReadBook(ctx context.Context, id int) (Book, error)
	ReadBooksForUser(ctx context.Context, userID uuid.UUID) ([]Book, error)
	SaveBook(ctx context.Context, book Book) error
	SaveBookMetadata(ctx context.Context, metadata BookMetadata) error
	SaveBookNotes(ctx context.Context, notes []BookNote) error
	SaveBookSubjects(ctx context.Context, subjects []BookSubject) error
	SaveUserBook(ctx context.Context, userBook UserBook) error
}
