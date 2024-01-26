package s3service

import (
	"backend/loggers"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// GetYears returns a list of years representing the available photoshoots in the S3 bucket.
// It retrieves the list of objects from the S3 bucket with the specified prefix and extracts the years from the common prefixes.
// The years are returned as a slice of strings, or error if the operation failed.
func (s *service) GetPhotoshootYears(ctx context.Context) ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := "photoshoots/"

	// call to list
	output, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from s3: %v", err)
	}

	var years []string
	for _, content := range output.CommonPrefixes {
		year := strings.TrimPrefix(*content.Prefix, prefix)
		year = strings.TrimSuffix(year, "/")
		years = append(years, year)
	}

	elapsedTime := time.Since(startTime)
	loggers.Performance.Printf("GetYears took %s", elapsedTime)
	return years, nil
}

// GetEvents retrieves a list of events for a given year from the S3 bucket.
// It takes a context.Context and a year string as input parameters.
// It returns a slice of strings representing the events and an error if any.
func (s *service) GetPhotoshootEvents(ctx context.Context, year string) ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("photoshoots/%s/", year)

	output, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from s3: %v", err)
	}

	var events []string
	for _, content := range output.CommonPrefixes {
		event := strings.TrimPrefix(*content.Prefix, prefix)
		event = strings.TrimSuffix(event, "/")
		events = append(events, event)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("GetEvents took %s", elapsedTime)

	return events, nil
}

// ListPhotos retrieves a list of photos from the S3 bucket for a specific year and event.
// It takes a context.Context, year, and event as parameters.
// It returns a slice of strings containing the photo names and an error if any.
func (s *service) ListPhotoshootPhotos(ctx context.Context, year, event string) ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("photoshoots/%s/%s/", year, event)

	output, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from s3: %v", err)
	}

	var photos []string
	for _, content := range output.Contents {
		photo := strings.TrimPrefix(*content.Key, prefix)
		photos = append(photos, photo)
	}

	elapsedTime := time.Since(startTime)
	loggers.Performance.Printf("ListPhotos took %s", elapsedTime)
	return photos, nil
}

// GetPhotos retrieves the URLs of photos from the specified S3 bucket for a given year and event.
// It takes a context.Context, year, and event as input parameters.
// It returns a slice of strings containing the photo URLs and an error if any.
func (s *service) GetPhotoshootPhotos(ctx context.Context, year, event string) ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("photoshoots/%s/%s/", year, event)

	output, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from s3: %v", err)
	}

	var photoURLs []string
	for _, content := range output.Contents {
		request, err := s.presigner.GetObject(ctx, bucket, *content.Key, 900) // 900 seconds = 15 minutes
		if err != nil {
			log.Printf("failed to create presigned URL for %s: %v", *content.Key, err)
			continue // log error and continue with the next object
		}

		photoURLs = append(photoURLs, request.URL)
	}

	elapsedTime := time.Since(startTime)
	loggers.Performance.Printf("GetPhotos took %s", elapsedTime)
	return photoURLs, nil
}
