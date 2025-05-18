package helper

import (
	"math/rand"
)

const (
	appUrl = "localhost:8080"
)

func ConvertShortCodeToShortUrl(shortCode string) string {
	return appUrl + "/" + shortCode
}

func CreateShortCodeKey() string {
	const allowedChars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	const length = 6

	shortCodeKey := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(len(allowedChars))
		shortCodeKey = shortCodeKey + string(allowedChars[index])
	}
	return shortCodeKey
}
