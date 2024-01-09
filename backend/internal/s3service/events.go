package s3service

import (
	"fmt"
	"log"
	"os"
	"time"
)

func (s *service) GenerateUploadURL(eventID, filename string, lifetimeSecs int64) (string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("events/%s/%s", eventID, filename)

	req, err := s.presigner.PutObject(bucket, prefix, lifetimeSecs)
	if err != nil {
		return "", fmt.Errorf("failed to get presigned url: %v", err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("GenerateUploaderURL took %s", elapsedTime)

	return req.URL, nil
}

func (s *service) DevGenerateUploadURL(eventID, filename string, lifetimeSecs int64) (string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("testing/%s/%s", eventID, filename)

	req, err := s.presigner.PutObject(bucket, prefix, lifetimeSecs)
	if err != nil {
		return "", fmt.Errorf("failed to get presigned url: %v", err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("GenerateUploaderURL took %s", elapsedTime)

	return req.URL, nil
}
