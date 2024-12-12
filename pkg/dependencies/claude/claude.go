package claude

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gutenberg/pkg"
	"io"
	"log"
	"net/http"
)

type claude struct {
	key string
}

func NewResponseGenerator(key string) pkg.ResponseGenerator {
	return &claude{
		key: key,
	}
}

func (c *claude) GenerateResponse(ctx context.Context, prompt string) (string, error) {
	messageReq := messageRequest{
		Model:       Claude35Sonnet,
		Temperature: 0,
		MaxTokens:   4000,
		System:      prompt,
		Messages: []message{
			{
				Role: "user",
				Content: []content{
					{
						Type: ContentTypeText,
						Text: prompt,
					},
				},
			},
		},
	}

	data, err := json.Marshal(messageReq)
	if err != nil {
		log.Println("error marshalling message request", err)
		return "", errors.New("error marshalling message request")
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.anthropic.com/v1/messages", bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Println("error creating request", err)
		return "", errors.New("error creating request")
	}

	req.Header.Set("x-api-key", c.key)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request", err)
		return "", errors.New("error creating request")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Println("unexpected status code", resp.StatusCode, string(body))
		return "", errors.New("unexpected status code")
	}

	var response messageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("error unmarshalling response", err)
		return "", errors.New("error unmarshalling response")
	}

	if len(response.Content) == 0 {
		log.Println("no content in response")
		return "", errors.New("no content in response")
	}

	return response.Content[0].Text, nil
}
