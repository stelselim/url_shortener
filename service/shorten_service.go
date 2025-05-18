package service

import (
	"context"
	"url_shortener/helper"
	"url_shortener/model"
)

func ShortenURL(originalUrl string) (string, error) {
	ctx := context.Background()
	docRef, _ := GetDocumentByOriginalURL(ctx, originalUrl)

	// If originalUrl exists, do not create new one.
	if docRef != nil {
		shortUrl, docErr := GetShortenedUrlModelByDocRef(ctx, docRef)

		if docErr != nil {
			return "", docErr
		}

		return helper.ConvertShortCodeToShortUrl(shortUrl.ShortCode), nil
	}

	// If does not exist, Create one doc for the original URL.
	shortCode := helper.CreateShortCodeKey()
	success, createErr := CreateShortenedUrlDocument(ctx, originalUrl, shortCode)

	if success && createErr != nil {
		return "", createErr
	}

	return helper.ConvertShortCodeToShortUrl(shortCode), nil
}

func GetOriginalUrl(shortcode string) (string, error) {
	var shortUrl model.ShortenedURL
	ctx := context.Background()
	docRef, docErr := GetDocumentByShortCode(ctx, shortcode)

	if docErr != nil {
		return "", docErr
	}
	shortUrl, err := GetShortenedUrlModelByDocRef(ctx, docRef)

	if err != nil {
		return "", err
	}

	return shortUrl.OriginalUrl, nil
}

func GetShortenedUrlStats(shortcode string) (model.ShortenedURL, error) {
	var shortUrl model.ShortenedURL
	ctx := context.Background()
	docRef, docErr := GetDocumentByShortCode(ctx, shortcode)

	if docErr != nil {
		return shortUrl, docErr
	}
	shortUrl, err := GetShortenedUrlModelByDocRef(ctx, docRef)

	if err != nil {
		return shortUrl, docErr
	}

	return shortUrl, nil
}

func DeleteShortenedUrl(shortcode string) error {
	ctx := context.Background()
	docRef, err := GetDocumentByShortCode(ctx, shortcode)
	if err != nil {
		return err
	}
	_, deleteErr := docRef.Delete(ctx)

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func IncreaseClicksByOriginalUrl(originalUrl string) (bool, error) {
	ctx := context.Background()
	docRef, err := GetDocumentByOriginalURL(ctx, originalUrl)
	if err != nil {
		return false, err
	}
	return IncreaseClickByOne(ctx, docRef)
}
