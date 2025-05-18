package service

import "url_shortener/model"

func ShortenURL() (string, error) {
	return "", nil
}

func GetOriginalUrl() (string, error) {
	return "", nil
}

func GetShortenedUrlStats() (model.ShortenedURL, error) {
	//TODO:
	return model.ShortenedURL{}, nil
}

func DeleteShortenedUrl() error {
	return nil
}
