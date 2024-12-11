package ebooks

import (
	"log"
	"testing"
)

func TestGetBook(t *testing.T) {
	t.Skip()
	client := NewClient()

	book, err := client.FetchBook(1)
	if err != nil {
		t.Error("error fetching book", err)
	}
	log.Println(book.Metadata.Subjects)
	t.Fatal(book)
}

func TestFetchBookText(t *testing.T) {
	t.Skip()
	client := NewClient()

	text, err := client.FetchBookText(1)
	if err != nil {
		t.Error("error fetching book text", err)
	}
	log.Println(text)
	t.Fatal(text)
}
