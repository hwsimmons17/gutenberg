package postgres

import (
	"context"
	"database/sql"
	"gutenberg/pkg"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type repository struct {
	db *bun.DB
}

func InitRepository(url string) pkg.BookRepository {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))
	maxOpenConns := 80
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)
	client := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	// client.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return &repository{
		db: client,
	}
}

func (r *repository) ReadBook(ctx context.Context, bookID int) (pkg.Book, error) {
	var book pkg.Book
	err := r.db.NewSelect().Model(&book).Where("book.id = ?", bookID).Relation("Metadata").Relation("Metadata.Notes").Relation("Metadata.Subjects").Scan(ctx)
	return book, err
}

func (r *repository) ReadBooksForUser(ctx context.Context, userID uuid.UUID) ([]pkg.Book, error) {
	var books []pkg.Book
	err := r.db.NewSelect().Model(&books).Join("JOIN user_books ON user_books.book_id = book.id").Where("user_books.user_id = ?", userID).Relation("Metadata").Relation("Metadata.Notes").Relation("Metadata.Subjects").Scan(ctx)
	return books, err
}

func (r *repository) SaveBook(ctx context.Context, book pkg.Book) error {
	_, err := r.db.NewInsert().Model(&book).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	return err
}

func (r *repository) SaveBookMetadata(ctx context.Context, metadata pkg.BookMetadata) error {
	_, err := r.db.NewInsert().Model(&metadata).On("CONFLICT (book_id) DO UPDATE").Exec(ctx)
	return err
}

func (r *repository) SaveBookNotes(ctx context.Context, notes []pkg.BookNote) error {
	_, err := r.db.NewInsert().Model(&notes).Exec(ctx)
	return err
}

func (r *repository) SaveBookSubjects(ctx context.Context, subjects []pkg.BookSubject) error {
	_, err := r.db.NewInsert().Model(&subjects).Exec(ctx)
	return err
}

func (r *repository) SaveUserBook(ctx context.Context, userBook pkg.UserBook) error {
	exists, err := r.db.NewSelect().
		Model(&userBook).
		Where("user_id = ? AND book_id = ?", userBook.UserID, userBook.BookID).
		Exists(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	_, err = r.db.NewInsert().Model(&userBook).Exec(ctx)
	return err
}
