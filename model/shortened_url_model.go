package model

type ShortenedURL struct {
	OriginalUrl string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	CreatedAt   string `json:"created_at"`
	Clicks      int    `json:"clicks"`
}
