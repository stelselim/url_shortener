package model

type ShortenPostResponseData struct {
	ShortUrl string `json:"short_url"`
}

type CreateShortenURLRequestData struct {
	OriginalUrl string `json:"original_url"`
}
