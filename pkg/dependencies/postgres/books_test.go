package postgres

import (
	"context"
	"gutenberg/pkg"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestReadBook(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	book, err := rep.ReadBook(context.Background(), 1)
	if err != nil {
		t.Error("error reading book", err)
	}
	t.Fatal(book)
}

func TestReadBooksForUser(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	books, err := rep.ReadBooksForUser(context.Background(), uuid.MustParse("0d99110d-2311-4ca1-b81a-f87ed73a1541"))
	if err != nil {
		t.Error("error reading books for user", err)
	}
	t.Fatal(books)
}

func TestSaveBook(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	if err := rep.SaveBook(context.Background(), pkg.Book{
		ID:       1,
		Title:    "Test Book",
		Author:   "Test Author",
		Metadata: &pkg.BookMetadata{},
	}); err != nil {
		t.Error("error saving book", err)
	}
}

func TestSaveBookMetadata(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	testText := "test"
	if err := rep.SaveBookMetadata(context.Background(), pkg.BookMetadata{
		BookID:              1,
		Language:            &testText,
		Summary:             &testText,
		Category:            &testText,
		ReleaseDate:         &testText,
		MostRecentlyUpdated: &testText,
		CopyrightStatus:     &testText,
		Downloads:           &testText,
		Notes:               []pkg.BookNote{},
		Subjects:            []pkg.BookSubject{},
	}); err != nil {
		t.Error("error saving book metadata", err)
	}
}

func TestSaveBookNotes(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	testText := "test"
	if err := rep.SaveBookNotes(context.Background(), []pkg.BookNote{
		{
			BookID: 1,
			Note:   testText,
		},
	}); err != nil {
		t.Error("error saving book notes", err)
	}
}

func TestSaveBookSubjects(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	testText := "test"
	if err := rep.SaveBookSubjects(context.Background(), []pkg.BookSubject{
		{
			BookID:  1,
			Subject: testText,
		},
	}); err != nil {
		t.Error("error saving book subjects", err)
	}
}

func TestSaveUserBook(t *testing.T) {
	t.Skip()
	godotenv.Load("../../../.env")
	url := os.Getenv("DATABASE_URL")
	rep := InitRepository(url)

	if err := rep.SaveUserBook(context.Background(), pkg.UserBook{
		UserID: uuid.New(),
		BookID: 1,
	}); err != nil {
		t.Error("error saving user book", err)
	}
}
