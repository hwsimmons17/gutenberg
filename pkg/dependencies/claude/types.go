package claude

type ImageMediaType string

const (
	ImageMediaTypeJpeg ImageMediaType = "image/jpeg"
	ImageMediaTypePng  ImageMediaType = "image/png"
	ImageMediaTypeGif  ImageMediaType = "image/gif"
	ImageMediaTypeWebp ImageMediaType = "image/webp"
)

type ContentType string

const (
	ContentTypeText  ContentType = "text"
	ContentTypeImage ContentType = "image"
)

type SourceType string

const (
	SourceTypeData SourceType = "base64"
)

type Model string

const (
	Claude3Opus    = "claude-3-opus-20240229"
	Claude3Sonnet  = "claude-3-sonnet-20240229"
	Claude3Haiku   = "claude-3-haiku-20240307"
	Claude35Sonnet = "claude-3-5-sonnet-20241022"
)

type messageRequest struct {
	Model       string    `json:"model"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
	System      string    `json:"system"`
	Messages    []message `json:"messages"`
}

type message struct {
	Role    string    `json:"role"`
	Content []content `json:"content"`
}

type content struct {
	Type   ContentType `json:"type"`
	Text   string      `json:"text,omitempty"`
	Source *source     `json:"source,omitempty"`
}

type source struct {
	Data      string         `json:"data"`
	MediaType ImageMediaType `json:"media_type"`
	Type      SourceType     `json:"type"`
}

type messageResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Role    string `json:"role"`
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	Model        string `json:"model"`
	StopReason   string `json:"stop_reason"`
	StopSequence any    `json:"stop_sequence"`
	Usage        struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}
