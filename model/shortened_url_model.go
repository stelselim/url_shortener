package model

type ShortenedURL struct {
	OriginalUrl string `json:"original_url" firestore:"original_url"`
	ShortCode   string `json:"short_code" firestore:"short_code"`
	CreatedAt   string `json:"created_at" firestore:"created_at"`
	Clicks      int    `json:"clicks" firestore:"clicks"`
}
