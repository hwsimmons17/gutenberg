package handlers

import (
	"context"
	"gutenberg/pkg"

	"github.com/google/uuid"
)

func GetBook(ctx context.Context, bookID int, userID uuid.UUID, bookRepo pkg.BookRepository, bookReader pkg.BookReader) (pkg.BookWithText, error) {
	book, err := bookRepo.ReadBook(ctx, bookID)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return pkg.BookWithText{}, err
	}
	text, errr := bookReader.FetchBookText(bookID)
	if errr != nil {
		return pkg.BookWithText{}, err
	}
	if err == nil {
		return pkg.BookWithText{Book: book, Text: text}, nil
	}
	newBook, err := bookReader.FetchBook(bookID)
	if err != nil {
		return pkg.BookWithText{}, err
	}
	if err := bookRepo.SaveBook(ctx, newBook); err != nil {
		return pkg.BookWithText{}, err
	}
	if err := bookRepo.SaveBookMetadata(ctx, *newBook.Metadata); err != nil {
		return pkg.BookWithText{}, err
	}
	if err := bookRepo.SaveBookNotes(ctx, newBook.Metadata.Notes); err != nil {
		return pkg.BookWithText{}, err
	}
	if err := bookRepo.SaveBookSubjects(ctx, newBook.Metadata.Subjects); err != nil {
		return pkg.BookWithText{}, err
	}
	book = newBook
	if err := bookRepo.SaveUserBook(ctx, pkg.UserBook{
		UserID: userID,
		BookID: bookID,
	}); err != nil {
		return pkg.BookWithText{}, err
	}

	return pkg.BookWithText{Book: book, Text: text}, nil
}

func GetBooks(ctx context.Context, userID uuid.UUID, bookRepo pkg.BookRepository) ([]pkg.Book, error) {
	return bookRepo.ReadBooksForUser(ctx, userID)
}
