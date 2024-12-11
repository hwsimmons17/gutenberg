package ebooks

import (
	"log"
	"testing"
)

func TestGetBook(t *testing.T) {
	client := NewClient()

	book, err := client.FetchBook(1)
	if err != nil {
		t.Error("error fetching book", err)
	}
	log.Println(book.Metadata.Subjects)
	t.Fatal(book)
}
