package ebooks

import (
	"errors"
	"fmt"
	"gutenberg/pkg"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type bookClient struct{}

func NewClient() pkg.BookReader {
	return &bookClient{}
}

func (c *bookClient) FetchBook(id int) (pkg.Book, error) {
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.gutenberg.org/ebooks/%d", id), nil)
	if err != nil {
		log.Println("error creating request", err)
		return pkg.Book{}, errors.New("error creating request")
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		log.Println("error sending request", err)
		return pkg.Book{}, errors.New("error sending request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("unexpected status code", resp.StatusCode)
		return pkg.Book{}, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		log.Println("error parsing html", err)
		return pkg.Book{}, errors.New("error parsing html")
	}

	return parseNodeToBook(id, node)
}

func parseNodeToBook(id int, node *html.Node) (pkg.Book, error) {
	book := pkg.Book{
		ID:     id,
		Title:  "",
		Author: "",
		Metadata: &pkg.BookMetadata{
			ID:                  id,
			BookID:              id,
			Language:            new(string),
			Summary:             new(string),
			Category:            new(string),
			ReleaseDate:         new(string),
			MostRecentlyUpdated: new(string),
			CopyrightStatus:     new(string),
			Downloads:           new(string),
			Notes:               []pkg.BookNote{},
			Subjects:            []pkg.BookSubject{},
		},
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, attr := range n.Attr {
				if attr.Key == "name" && attr.Val == "title" {
					for _, attr := range n.Attr {
						if attr.Key == "content" {
							book.Title = attr.Val
							return
						}
					}
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "tr" {
			var isAuthor, isLanguage, isSummary, isCategory, isReleaseDate, isMostRecentlyUpdated, isCopyrightStatus, isDownloads, isNote, isSubject bool
			for _, attr := range n.Attr {
				if attr.Key == "property" && attr.Val == "dcterms:language" {
					isLanguage = true
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Author" {
					isAuthor = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Summary" {
					isSummary = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Category" {
					isCategory = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Release Date" {
					isReleaseDate = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Most Recently Updated" {
					isMostRecentlyUpdated = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Copyright Status" {
					isCopyrightStatus = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Downloads" {
					isDownloads = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Note" {
					isNote = true
				}
				if c.Type == html.ElementNode && c.Data == "th" && c.FirstChild != nil && c.FirstChild.Data == "Subject" {
					isSubject = true
				}
				if isAuthor && c.Type == html.ElementNode && c.Data == "td" {
					for a := c.FirstChild; a != nil; a = a.NextSibling {
						if a.Type == html.ElementNode && a.Data == "a" {
							book.Author = a.FirstChild.Data
							return
						}
					}
				}
				if isSummary && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.Summary = &c.FirstChild.Data
					return
				}
				if isLanguage && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.Language = &c.FirstChild.Data
					return
				}
				if isCategory && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.Category = &c.FirstChild.Data
					return
				}
				if isReleaseDate && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.ReleaseDate = &c.FirstChild.Data
					return
				}
				if isMostRecentlyUpdated && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.MostRecentlyUpdated = &c.FirstChild.Data
					return
				}
				if isCopyrightStatus && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.CopyrightStatus = &c.FirstChild.Data
					return
				}
				if isDownloads && c.Type == html.ElementNode && c.Data == "td" {
					book.Metadata.Downloads = &c.FirstChild.Data
					return
				}
				if isNote && c.Type == html.ElementNode && c.Data == "td" {
					noteText := c.FirstChild.Data
					for a := c.FirstChild; a != nil; a = a.NextSibling {
						if a.Type == html.ElementNode && a.Data == "a" {
							noteText = noteText + a.FirstChild.Data
							break
						}
					}

					book.Metadata.Notes = append(book.Metadata.Notes, pkg.BookNote{
						BookID: id,
						Note:   noteText,
					})
					return
				}
				if isSubject && c.Type == html.ElementNode && c.Data == "td" {
					var subjectText string
					for a := c.FirstChild; a != nil; a = a.NextSibling {
						if a.Type == html.ElementNode && a.Data == "a" {
							subjectText = a.FirstChild.Data
							break
						}
					}
					book.Metadata.Subjects = append(book.Metadata.Subjects, pkg.BookSubject{
						BookID:  id,
						Subject: subjectText,
					})
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(node)

	if book.Title == "" {
		return pkg.Book{}, errors.New("title not found")
	}

	return book, nil
}

func (c *bookClient) FetchBookText(id int) (string, error) {
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.gutenberg.org/files/%d/%d-0.txt", id, id), nil)
	if err != nil {
		log.Println("error creating request", err)
		return "", errors.New("error creating request")
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		log.Println("error sending request", err)
		return "", errors.New("error sending request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("unexpected status code", resp.StatusCode)
		return "", fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return "", errors.New("error reading response body")
	}

	return string(body), nil
}
