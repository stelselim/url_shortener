package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"url_shortener/model"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	clientInstance *firestore.Client
	initOnce       sync.Once
	initErr        error
)

const (
	shortenedUrlsCollectionName = "shortenedUrls"
)

func GetFirestoreClient(ctx context.Context, options ...option.ClientOption) (*firestore.Client, error) {
	initOnce.Do(func() {
		app, err := firebase.NewApp(ctx, nil, options...)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize firebase app: %w", err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize firebase app: %w", err)
		}
		log.Println("firebase successfully initiliazed.")
		clientInstance = client
	})

	return clientInstance, initErr
}

func CloseFirestoreClient() error {
	if clientInstance != nil {
		return clientInstance.Close()
	}
	return nil
}

func GetDocumentByOriginalURL(ctx context.Context, originalUrl string) (*firestore.DocumentRef, error) {
	client, err := GetFirestoreClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("firestore client not initialized: %w", err)
	}

	docs, err := client.Collection(shortenedUrlsCollectionName).
		Where("original_url", "==", originalUrl).
		Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("error querying documents: %w", err)
	}

	if len(docs) == 0 {
		return nil, fmt.Errorf("no docs found: %w", err)
	}

	if len(docs) > 1 {
		// Delete other docs.
		for i := 1; i < len(docs); i++ {
			docs[i].Ref.Delete(ctx)
		}
	}

	if !docs[0].Exists() {
		return nil, fmt.Errorf("the doc not found")
	}

	return docs[0].Ref, nil
}

func GetDocumentByShortCode(ctx context.Context, shortCode string) (*firestore.DocumentRef, error) {
	client, err := GetFirestoreClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("firestore client not initialized: %w", err)
	}
	docs, docErr := client.Collection(shortenedUrlsCollectionName).Where("short_code", "==", shortCode).Documents(ctx).GetAll()

	if docErr != nil {
		return nil, fmt.Errorf("this shortcode has not been added: %w", docErr)
	}

	if len(docs) == 0 {
		return nil, fmt.Errorf("no docs found: %w", err)
	}

	if len(docs) > 1 {
		// Delete other docs.
		for i := 1; i < len(docs); i++ {
			docs[i].Ref.Delete(ctx)
		}
	}

	if !docs[0].Exists() {
		return nil, fmt.Errorf("the doc not found")
	}

	return docs[0].Ref, nil
}

func CreateShortenedUrlDocument(ctx context.Context, originalUrl string, shortCode string) (bool, error) {
	client, err := GetFirestoreClient(ctx)
	if err != nil {
		return false, fmt.Errorf("firestore client not initialized: %w", err)
	}

	newShortUrl := model.ShortenedURL{
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
		CreatedAt:   time.Now().Format(time.RFC3339), // More standard format
		Clicks:      0,
	}

	docRef, _, err := client.Collection(shortenedUrlsCollectionName).Add(ctx, newShortUrl)
	if err != nil {
		return false, fmt.Errorf("the shortened url could not created: %w", err)
	}
	fmt.Println("doc created at ID:", docRef.ID)
	return true, nil
}

func GetShortenedUrlModelByDocRef(ctx context.Context, docRef *firestore.DocumentRef) (model.ShortenedURL, error) {
	var shortUrl model.ShortenedURL

	docSnap, err := docRef.Get(ctx)
	if err != nil {
		return shortUrl, fmt.Errorf("error getting document by its ref: %s", err)
	}

	err = docSnap.DataTo(&shortUrl)
	if err != nil {
		return shortUrl, fmt.Errorf("error decoding document data:%s", err)
	}

	return shortUrl, nil
}

func IncreaseClickByOne(ctx context.Context, docRef *firestore.DocumentRef) (bool, error) {
	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "clicks", Value: firestore.Increment(1)},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
