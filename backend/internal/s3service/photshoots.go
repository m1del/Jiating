package s3service

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func (s *service) GetYears() ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := "photoshoots/"

	// call to list
	output, err := s.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
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
	log.Printf("GetYears took %s", elapsedTime)

	return years, nil
}

func (s *service) GetEvents(year string) ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("photoshoots/%s/", year)

	output, err := s.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
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

func (s *service) ListPhotos(year, event string) ([]string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("photoshoots/%s/%s/", year, event)

	output, err := s.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
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
	log.Printf("ListPhotos took %s", elapsedTime)

	return photos, nil
}

func (s *service) GetPhotos(ctx context.Context, year, event string) ([]string, error) {
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
		request, err := s.presigner.GetObject(bucket, *content.Key, 900) // 900 seconds = 15 minutes
		if err != nil {
			log.Printf("failed to create presigned URL for %s: %v", *content.Key, err)
			continue // log error and continue with the next object
		}

		photoURLs = append(photoURLs, request.URL)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("GetPhotos took %s", elapsedTime)

	return photoURLs, nil
}
