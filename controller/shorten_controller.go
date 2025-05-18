package controller

import (
	"fmt"
	"net/http"
	"url_shortener/helper"
	"url_shortener/model"
	"url_shortener/service"

	"github.com/labstack/echo/v4"
)

func PostShortenController(c echo.Context) error {
	shortUrl, err := service.ShortenURL()

	if err != nil {
		return helper.RespondError(
			c,
			http.StatusCreated,
			"URL could not be shortened.",
		)
	}

	return helper.RespondSuccess(
		c,
		http.StatusCreated,
		"URL was successfully shortened.",
		model.ShortenPostResponseData{
			ShortUrl: shortUrl,
		},
	)
}

func GetShortenCodeController(c echo.Context) error {
	originalUrl, err := service.GetOriginalUrl()

	if err != nil {
		return helper.RespondError(
			c,
			http.StatusCreated,
			"Original URL not found.",
		)
	}

	fmt.Printf("Redirecting to %s...\n", originalUrl)
	return c.Redirect(303, originalUrl)
}

func GetShortenCodeStatsController(c echo.Context) error {
	stats, err := service.GetShortenedUrlStats()
	if err != nil {
		return helper.RespondError(
			c,
			http.StatusCreated,
			"Stats not found",
		)
	}

	return helper.RespondSuccess(
		c,
		http.StatusCreated,
		"Short URL deleted successfully",
		stats,
	)
}

func DeleteShortenCodeController(c echo.Context) error {

	err := service.DeleteShortenedUrl()
	if err != nil {
		return helper.RespondError(
			c,
			http.StatusCreated,
			"URL Could not deleted",
		)
	}

	return helper.RespondSuccess[any](
		c,
		http.StatusCreated,
		"Short URL deleted successfully",
		nil,
	)
}
