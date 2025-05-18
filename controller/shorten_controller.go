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
	var createShortenUrlRequestData model.CreateShortenURLRequestData

	if err := c.Bind(&createShortenUrlRequestData); err != nil {
		return helper.RespondError(
			c,
			http.StatusBadRequest,
			"invalid request body.",
		)
	}

	shortUrl, err := service.ShortenURL(createShortenUrlRequestData.OriginalUrl)
	if err != nil {
		return helper.RespondError(
			c,
			http.StatusInternalServerError,
			err.Error(),
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

func GetOriginalUrlController(c echo.Context) error {
	shortCode := c.Param("shortCode")
	fmt.Println("Short Code to check the stats:", shortCode)

	if shortCode == "" {
		return helper.RespondError(
			c,
			http.StatusBadRequest,
			"Please, provide a shortcode",
		)
	}

	originalUrl, err := service.GetOriginalUrl(shortCode)
	if err != nil {
		return helper.RespondError(
			c,
			http.StatusNotFound,
			err.Error(),
		)
	}

	fmt.Printf("Redirecting to %s...\n", originalUrl)
	// Increase the clicks for URL.
	service.IncreaseClicksByOriginalUrl(originalUrl)

	return c.Redirect(http.StatusPermanentRedirect, originalUrl)
}

func GetShortenCodeStatsController(c echo.Context) error {
	shortCode := c.Param("shortCode")
	fmt.Println("Short Code to check the stats:", shortCode)

	if shortCode == "" {
		return helper.RespondError(
			c,
			http.StatusBadRequest,
			"Please, provide a shortcode",
		)
	}

	stats, err := service.GetShortenedUrlStats(shortCode)
	if err != nil {
		return helper.RespondError(
			c,
			http.StatusNotFound,
			err.Error(),
		)
	}

	return helper.RespondSuccess(
		c,
		http.StatusOK,
		"Stats found successfully",
		stats,
	)
}

func DeleteShortenCodeController(c echo.Context) error {
	shortCode := c.Param("shortCode")
	fmt.Println("Short Code to be deleted:", shortCode)

	if shortCode == "" {
		return helper.RespondError(
			c,
			http.StatusBadRequest,
			"Please, provide a shortcode",
		)
	}

	err := service.DeleteShortenedUrl(shortCode)
	if err != nil {
		return helper.RespondError(
			c,
			http.StatusNotFound,
			err.Error(),
		)
	}

	return helper.RespondSuccess[any](
		c,
		http.StatusOK,
		"Short URL deleted successfully",
		nil,
	)
}
